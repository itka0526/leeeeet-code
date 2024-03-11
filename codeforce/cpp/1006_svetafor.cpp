#include <bits/stdc++.h>

using namespace std;

int main()
{
    string a, b, c;
    cin >> a >> b >> c;
    if ((a == "black") && (b == "black") && (c == "green"))
    {
        cout << "black" << endl << "black" << endl << "GREEN";
    }
    else if ((a == "black") && (b == "black") && (c == "GREEN"))
    {
        cout << "black" << endl << "yellow" << endl << "black";
    }
    else if ((a == "black") && (b == "yellow") && (c == "black"))
    {
        cout << "red" << endl << "black" << endl << "black";
    }
    else if ((a == "red") && (b == "black") && (c == "black"))
    {
        cout << "red" << endl << "yellow" << endl << "black";
    }
    else if ((a == "red") && (b == "yellow") && (c == "black"))
    {
        cout << "black" << endl << "black" << endl << "green";
    }
    else if ((a == "black") && (b == "YELLOW") && (c == "black"))
    {
        cout << "black" << endl << "YELLOW" << endl << "black";
    }
    else
    {
        cout << "error";
    }
    return 0;
}