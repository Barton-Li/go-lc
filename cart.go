package main

import (
    "context"
    "encoding/json"
    "fmt"
    "strconv"
    "sync"

    "github.com/go-redis/redis/v8"
)

type UserInfoTo struct {
    UserId   *int64
    UserKey  string
}

type CartItem struct {
    Check     bool
    Count     int
    Image     string
    Title     string
    Price     float64
    SkuId     int64
    SkuAttr   []string
}

type Cart struct {
    Items []CartItem
}

type SkuInfoVo struct {
    SkuDefaultImg string
    SkuTitle      string
    Price         float64
    SkuId         int64
}

type ShoppingService interface {
    GetSkuInfo(ctx context.Context, skuId int64) (R, error)
    GetSkuSaleAttrValues(ctx context.Context, skuId int64) ([]string, error)
}

type R map[string]interface{}

type CartService struct {
    RedisClient  *redis.Client
    ProductService ShoppingService
    Executor     sync.WaitGroup
    CART_PREFIX  string
}

func (c *CartService) getCartOps(userInfoTo UserInfoTo) *redis.HashCommands {
    var cartKey string
    if userInfoTo.UserId != nil {
        cartKey = c.CART_PREFIX + strconv.FormatInt(*userInfoTo.UserId, 10)
    } else {
        cartKey = c.CART_PREFIX + userInfoTo.UserKey
    }
    ctx := context.Background()
    return c.RedisClient.HSet(ctx)
}

func (c *CartService) addToCart(skuId int64, num int) (CartItem, error) {
    var cartItem CartItem

    ctx := context.Background()
    cartOps := c.getCartOps(CartInterceptor.threadLocal.Get())

    res, err := cartOps.Get(ctx, strconv.FormatInt(skuId, 10)).Result()
    if err == redis.Nil || res == "" {
        c.Executor.Add(2)

        // 商品添加到购物车(新商品)
        go func() {
            defer c.Executor.Done()
            skuInfo, err := c.ProductService.GetSkuInfo(ctx, skuId)
            if err != nil {
                fmt.Println(err)
                return
            }

            data := skuInfo["skuInfo"].(SkuInfoVo)
            cartItem.Check = true
            cartItem.Count = 1
            cartItem.Image = data.SkuDefaultImg
            cartItem.Title = data.SkuTitle
            cartItem.Price = data.Price
            cartItem.SkuId = data.SkuId
        }()

        // 运程查询sku的组合信息
        go func() {
            defer c.Executor.Done()
            values, err := c.ProductService.GetSkuSaleAttrValues(ctx, skuId)
            if err != nil {
                fmt.Println(err)
                return
            }
            cartItem.SkuAttr = values
        }()

        c.Executor.Wait()

        byteData, err := json.Marshal(cartItem)
        if err != nil {
            return cartItem, err
        }

        cartOps.Set(ctx, strconv.FormatInt(skuId, 10), string(byteData))

        return cartItem, nil
    } else {
        if err := json.Unmarshal([]byte(res), &cartItem); err != nil {
            return cartItem, err
        }
        cartItem.Count += num
        byteData, err := json.Marshal(cartItem)
        if err != nil {
            return cartItem, err
        }

        cartOps.Set(ctx, strconv.FormatInt(skuId, 10), string(byteData))

        return cartItem, nil
    }
}

func (c *CartService) getCartItem(skuId int64) (CartItem, error) {
    var cartItem CartItem

    ctx := context.Background()
    cartOps := c.getCartOps(CartInterceptor.threadLocal.Get())

    res, err := cartOps.Get(ctx, strconv.FormatInt(skuId, 10)).Result()
    if err == redis.Nil || res == "" {
        return cartItem, fmt.Errorf("item not found")
    }

    if err := json.Unmarshal([]byte(res), &cartItem); err != nil {
        return cartItem, err
    }

    return cartItem, nil
}

func (c *CartService) getCartItems(cartKey string) ([]CartItem, error) {
    ctx := context.Background()
    cartOps := c.RedisClient.HGetAll(ctx, cartKey)

    values, err := cartOps.Result()
    if err != nil {
        return nil, err
    }

    var cartItems []CartItem
    for _, v := range values {
        var item CartItem
        if err := json.Unmarshal([]byte(v), &item); err != nil {
            return nil, err
        }
        cartItems = append(cartItems, item)
    }

    return cartItems, nil
}

func (c *CartService) getCart() (Cart, error) {
    var cart Cart

    userInfoTo := CartInterceptor.threadLocal.Get()
    ctx := context.Background()

    if userInfoTo.UserId != nil {
        // 登录
        cartKey := c.CART_PREFIX + strconv.FormatInt(*userInfoTo.UserId, 10)
        tempCartKey := c.CART_PREFIX + userInfoTo.UserKey

        // 如果临时购物车的数据还没有合并【合并购物车】
        tempCartItems, err := c.getCartItems(tempCartKey)
        if err != nil {
            return cart, err
        }

        if tempCartItems != nil {
            // 合并
            for _, item := range tempCartItems {
                _, _ = c.addToCart(item.SkuId, item.Count)
            }
        }

        // 获取登录后的数据
        cartItems, err := c.getCartItems(cartKey)
        if err != nil {
            return cart, err
        }

        cart.Items = cartItems
        // 清除临时购物车数据
        c.clearCart(tempCartKey)
    } else {
        // 没登陆
        cartKey := c.CART_PREFIX + userInfoTo.UserKey
        // 获取临时购物车的所有购物项
        cartItems, err := c.getCartItems(cartKey)
        if err != nil {
            return cart, err
        }

        cart.Items = cartItems
    }

    return cart, nil
}

func (c *CartService) checkItem(skuId int64, check int) error {
    cartOps := c.getCartOps(CartInterceptor.threadLocal.Get())

    cartItem, err := c.getCartItem(skuId)
    if err != nil {
        return err
    }

    cartItem.Check = check == 1
    byteData, err := json.Marshal(cartItem)
    if err != nil {
        return err
    }

    return cartOps.Set(context.Background(), strconv.FormatInt(skuId, 10), string(byteData)).Err()
}

func (c *CartService) changeItemCount(skuId int64, num int) error {
    cartOps := c.getCartOps(CartInterceptor.threadLocal.Get())

    cartItem, err := c.getCartItem(skuId)
    if err != nil {
        return err
    }

    cartItem.Count = num
    byteData, err := json.Marshal(cartItem)
    if err != nil {
        return err
    }

    return cartOps.Set(context.Background(), strconv.FormatInt(skuId, 10), string(byteData)).Err()
}

func (c *CartService) deleteItem(skuId int64) error {
    cartOps := c.getCartOps(CartInterceptor.threadLocal.Get())

    return cartOps.Del(context.Background(), strconv.FormatInt(skuId, 10)).Err()
}

func (c *CartService) clearCart(cartKey string) error {
    return c.RedisClient.Del(context.Background(), cartKey).Err()
}

type ThreadLocal struct {
    sync.Map
}

var CartInterceptor = struct {
    threadLocal ThreadLocal
}{}

func (t *ThreadLocal) Get() UserInfoTo {
    val, _ := t.Load("userInfo")
    return val.(UserInfoTo)
}

func (t *ThreadLocal) Set(userInfo UserInfoTo) {
    t.Store("userInfo", userInfo)
}

func main() {
    // 初始化Redis客户端和服务
    redisClient := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // 如果没有密码则为空
        DB:       0,  // 默认数据库
    })

    productService := &ProductServiceImpl{}

    cartService := &CartService{
        RedisClient:  redisClient,
        ProductService: productService,
        CART_PREFIX:  "cart:",
    }

    // 设置线程本地存储
    CartInterceptor.threadLocal.Set(UserInfoTo{UserId: new(int64)})
    *CartInterceptor.threadLocal.Get().UserId = 1

    // 添加商品到购物车
    cartItem, err := cartService.addToCart(100, 1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("CartItem added: %+v\n", cartItem)
    }

    // 获取购物车
    cart, err := cartService.getCart()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Cart: %+v\n", cart)
    }
}

// 假设的远程服务实现
type ProductServiceImpl struct{}

func (p *ProductServiceImpl) GetSkuInfo(ctx context.Context, skuId int64) (R, error) {
    // 模拟远程服务返回
    return R{
        "skuInfo": SkuInfoVo{
            SkuDefaultImg: "default.jpg",
            SkuTitle:      "商品标题",
            Price:         100.0,
            SkuId:         skuId,
        },
    }, nil
}

func (p *ProductServiceImpl) GetSkuSaleAttrValues(ctx context.Context, skuId int64) ([]string, error) {
    // 模拟远程服务返回
    return []string{"attr1", "attr2"}, nil
}
