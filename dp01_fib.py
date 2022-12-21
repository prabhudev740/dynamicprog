def fib(n):
    if n <= 2:
        return 1
    return fib(n-1) + fib(n-2)


def fib_dy(n, data={}):
    if n <= 2:
        return 1
    if n in data:
        return data[n]
    res = fib_dy(n-1, data) + fib_dy(n-2, data)
    data[n] = res
    return res

def main():
    data = {}
    print(fib_dy(3))
    print(fib_dy(4))
    print(fib_dy(5))
    print(fib_dy(100) )

main()

# 					5
#  			  4				3
#  		  3	     2	 	  2	   1
#       2	1