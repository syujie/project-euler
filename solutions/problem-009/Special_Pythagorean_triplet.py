a = 1
while a <= int(1000/3):
	b = a+1
	while b <= 500:
		c = 1000 - a - b
		if a**2 + b**2 == c**2:
			print(a*b*c)
			break
		b += 1
	a += 1