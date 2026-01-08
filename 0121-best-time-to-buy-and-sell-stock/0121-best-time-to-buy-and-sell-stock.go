func maxProfit(prices []int) int {
    minPrice := int(^uint(0) >> 1) 
    maxProfit := 0

    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        } else if price-minPrice > maxProfit {
            maxProfit = price - minPrice
        }
    }

    return maxProfit

    
}