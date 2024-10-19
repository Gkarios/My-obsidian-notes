$$d(x,y)=||x-y||_{2}=\sqrt{\sum _{i=1}^{d}(x_{i}-y_{i})^{2}}.$$
The upper limit of the summation notation is the d=dimension (2D, 3D). 

distance of x,y = L2 norm of the |x-y| vector.
This equals the square root of summation from 1 to d dimensions of ($x_{i} - y_{i}$). 

The i index of summation, is the current coordinate. 

So in a 2D Cartesian coordinate system, when it comes to a point x in space $x_{1}$ is the abscissa and $x_{2}$ is the ordinate. 

$$||x-y||_{2}=\sqrt{\sum _{i=1}^{d}(x_{i}-y_{i})^{2}}$$
2 dimensions, so d = 2: 
$(x_{1} - y_{1})^{2} + (x_{2} + y_{2})^{2}$

$(\frac{N(N-1)}{2})$ checks and a time complexity of $(\Theta(N^2))$ due to the following reasons:
$$ [  
(N-1) + (N-2) + \ldots + 1 + 0 = \frac{N(N-1)}{2}  
]  
This is the sum of the first ( N-1 ) natural numbers.$$

- he sum of the first ( k ) natural numbers is given by the formula:  
    $$[  
    \sum_{i=1}^{k} i = \frac{k(k+1)}{2}  
    ]
- Here, ( k = N-1 ). Therefore:  
    [  
    \sum_{i=1}^{N-1} i = \frac{(N-1)N}{2}  
    ]$$