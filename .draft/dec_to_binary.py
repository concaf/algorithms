def d2b(decimal=17):
    binary = ""
    while decimal >= 1:
        binary += str(decimal % 2)
        decimal /= 2
    return binary

print(d2b())