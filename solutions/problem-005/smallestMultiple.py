def smallestMultiple(n):
	numberToCheck = n
	found = False

	while not found:
		if check(numberToCheck, n):
			found = True
		numberToCheck += n
	return numberToCheck - n

def check(numberToCheck, number):
	for i in range(1, number+1):
		if numberToCheck%i != 0:
			return False
	return True
print(smallestMultiple(10))