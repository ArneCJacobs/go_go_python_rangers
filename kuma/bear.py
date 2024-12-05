import math
import random


def detect(data: list[float]) -> list[float]:
    print("Hello from python")
    data_new = [i + 1 for i in data]
    return data_new
     


def gen_test_data() -> list[float]:
    return [
        (random.random() * 2.0 - 1.0) * 100.0
        for _ in range(10)
    ]
