import random

# generate 30 random bitcoin prices
def random_prices():
    prices = [random.uniform(2081.36, 2082.36) for _ in range(30)]
    print(prices)
    
random_prices()