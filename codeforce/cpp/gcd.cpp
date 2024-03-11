#include <bits/stdc++.h>

using namespace std;

void writeln()
{
    cout << endl;
}

int sqr(int x)
{
    return x * x;
}

int sum(int a, int b, int c)
{
    return a + b + c;
}

int F(int n)
{
    int s = 1;
    for (int i = 2; i <= n; i++)
    {
        s *= i;
    }
    return s;
}

double Ts(double a, double b, double c)
{
    double p;
    p = (a + b + c) / 2;
    return sqrt(p * (p - a) * (p - b) * (p - c));
}

int gcd(int a, int b)
{
    while (b)
    {
        a %= b;
        swap(a, b);
    }
    return a;
}

int lcm(int a, int b)
{
    return a / gcd(a, b) * b;
}

int main()
{
    cout << sqr(2 + 3);
    writeln();
    cout << sum(1, 3, 5) << endl;
    cout << F(F(3) - 1) << endl;
    cout << Ts(3, 4, 5) << endl;
    cout << gcd(12, 16) << endl;
    cout << lcm(12, 16) << endl;
    return 0;
}