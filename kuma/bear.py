import math
import random
import numpy as np


def detect(data: list[float]) -> list[float]:
    print("Hello from python")
    data_new = [i + 1 for i in data]
    return data_new
     


def gen_test_data() -> list[float]:
    data = [
        (random.random() * 2.0 - 1.0) * 100.0
        for _ in range(10)
    ]
    print(f"hello from python, again: {data=}")
    return data
