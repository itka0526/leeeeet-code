#include <bits/stdc++.h>

using namespace std;

int main()
{
    int b1[3], b2[3];
    for (int i = 0; i < 3; i++)
        cin >> b1[i];
    for (int i = 0; i < 3; i++)
        cin >> b2[i];
    sort(b1, b1 + 3);
    sort(b2, b2 + 3);
    if (b1[0] == b2[0] && b1[1] == b2[1] && b1[2] == b2[2])
        cout << "Boxes are equal";
    else if (!(b1[0] > b2[0]) && !(b1[1] > b2[1]) && !(b1[2] > b2[2]))
        cout << "The first box is smaller than the second one";
    else if (!(b1[0] < b2[0]) && !(b1[1] < b2[1]) && !(b1[2] < b2[2]))
        cout << "The first box is larger than the second one";
    else
        cout << "Boxes are incomparable";
    return 0;
}