from datetime import datetime


def log_it(func):
    def wrapper():
        print("Befor Function Call")
        func()
        print("After Function Call")

    return wrapper


@log_it
def print_time():
    # Logic
    print(datetime.now())
    # Logic


print_time()
