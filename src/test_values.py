import random
from datetime import datetime, timedelta

# generate 30 random ethereum prices in a list
def random_prices_list():
    prices = [random.uniform(2000, 2100) for _ in range(30)]
    print(prices)
    
random_prices_list()

# generate n random ethereum prices into a dictionary where prices are mapped to a time
def random_timeseries_prices(n):
    prices = {}
    timestamp_begin = datetime.today() - timedelta(hours=0, minutes=n)
    for i in range(n):
        prices[(timestamp_begin + timedelta(minutes=i))] = random.uniform(2000, 2100)

    print(prices)
    return prices

def calculate_EMA(time_begin, prices):
    smoothing_factor = 2 / (1 + len(prices))
    current_EMA = prices[time_begin]
    for i in range(1, len(prices)):
        current_EMA = (prices[time_begin + timedelta(minutes=i)] * smoothing_factor) +  (current_EMA * (1 - smoothing_factor))
    print(current_EMA)
    
prices = random_timeseries_prices(60)
time_begin = next(iter(prices))

calculate_EMA(time_begin, prices)