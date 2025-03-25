import random

# generate 30 random bitcoin prices
def random_prices():
    prices = [random.uniform(84149.64, 84150.64) for _ in range(30)]
    print(prices)
    
random_prices()