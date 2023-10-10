import random

languages = {
    "important": ["Go", "C++", "C", "Python"],
    "medium": ["JavaScript", "Swift", "Kotlin"],
    "esoteric": ["Racket", "x86_64"]
}

random_number = random.randint(0, 100)
if random_number < 60:
    importance = "important"
elif random_number < 90:
    importance = "medium"
else:
    importance = "esoteric"

todays_language = random.choice(languages[importance])

print(f"Today, you will use {todays_language}!")