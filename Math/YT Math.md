# Math 4 everyone 
Set
Subset
Genuine Subset

Equation is an equality that is true for one value. Tautology is an equality that is true for every value.

Equation is an equality between two algebraic equations.

For every sets A, B:
$B\subseteq A$ ^ $A \subseteq B$ <==> A = B

f(x) = y
f: $\mathbb{R} \to \mathbb{R}$ (if x,y $\in \mathbb{R}$)
x: independent variable
y: dependent variable

Domain: all possible x values (ΠΟ)
Range: all possible y values (ΣΤ)

Monotonous function: A function that is consistently analogous or reverse-analogous throughout.

$$
\left\{ \begin{aligned} dy &= y - y_0 > 0 \\ dx &= x - x_0 > 0 \end{aligned} \right.  \implies \frac{dy}{dx} > 0 $$

- When are 2 values analogous? 
When their quotient is a parameter(constant). Which means that their function is linear and f(0)=0. 
$\frac{y}{x} = λ$

- What is a graph of a function?
The graph of a function is the set of points, the ordered pairs of the plane (usually the Cartesian plane {(x,y)$\in \mathbb{R} ^2$ , where y = f(x)} ). 
# Statistics 4 everyone
Remember this graph ALWAYS:
$[]^{Ω}\to_{RV}[]^{R}\to_{PD}[]^{p}$

Contingency: the result of random experiment.

Elemental contingency / sample point: One outcome of a random experiment (a 3 from a die roll).

Complex contingency: One outcome resulted from many elemental contingencies (even numbers from a die roll).

Sample space: A set of all the sample points that can be resulted from the random experiment.

Probability of an event: $\frac{W}{A}$, W: outcomes I want, A: all outcomes

Complementary contingency: $\overline{A} =  A - \frac{W}{A}$

Population: The whole set we study in regards to a characteristic

Sample: A portion of the population that interests us

- Conditional Probability
P(A/B) = $\frac{P(A\cap B)}{P(B)}$
<=> $P(A\cap B)$ = P(B) P(A/B) $^{(1)}$

P(B/A) = $\frac{P(B \cap A)}{P(A)}$
<=> $P(A \cap B)$ = P(A) P(B/A) $^{(2)}$

From $_{(1),(2)}$ : 
P(A) P(B/A) = P(B) P (A/B)
-> Bayes' law.

To test if two contingencies A, B are independent:
P(A$\cap$B) == P(A) P(B)
If not equal, the contingencies are dependent on one another.

- What is the criteria for when a random variable is discrete or continuous?
It is the range of the random variable. If its range is consisted of discrete numbers, the random variable is discrete.

Features of probability discrete functions:
i) $P_{x}(x) \geq 0$

ii) $\sum_{x=1}^{Ω}{P_{x}(x) = 1}$, Ω: range of RV

iii) $P(X\in B) =\sum_{x=1}^{B}P_{x}(x=i)$

Features of probability density functions:
i) $f_{x}(x)\geq 0$

ii) $\int f_{x}(x)dx=1$

- Cumulative distribution function (cdf)
For discrete functions:
$F_{x}(x)=P(X\geq x)$

P(α < X < β) = $F_{x}(β) - F_{x}(α)$

for density functions:
$F_{x}(x) = \int_{-\infty}^{x}f_{x}(y)dy$


# Euler's Formula
[By Eddie Woo](https://youtu.be/pg827uDPFqA?feature=shared)

Calculating interest:
A = P(1+n)^n

In the annual context:

100% interest rate per year would be:
A = (1+1)^1
=2

50% interest rate twice a year would be:
A = (1+1/2)^2
= 1+2 * 1/2+ 1/4 = 2.25

25% interest rate 4 times a year:
A = (1+1/4)^4
= 2.434

Every month: 100/12% interest rate 12 times a year:
A = (1+1/12)^12
= 2.61

Every day: 100/365% interest rate rate 365 times a year:
A = (1+1/365)^365
= 2.714

e = (1+infinite)^infinite

Binomial Theorem: 
![[Binomial Theorem.png]]