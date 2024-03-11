#include <bits/stdc++.h>

using namespace std;
int main()
{
    int l1, w1, h1, l2, w2, h2, lc, wc, hc, l;
    bool ind = true;
    cin >> l1 >> w1 >> h1 >> l2 >> w2 >> h2 >> lc >> wc >> hc;
    l = l1;
    l1 = max(l1, w1);
    w1 = min(l, w1);
    l = l2;
    l2 = max(l2, w2);
    w2 = min(l, w2);
    l = lc;
    lc = max(lc, wc);
    wc = min(l, wc);
    if (l1 <= lc && w1 <= wc && l2 <= lc && w2 <= wc)
    {
        if (h1 > hc || h2 > hc)
        {
            printf("NO");
            ind = false;
        }
        if (ind && ind && l1 <= lc && w1 <= wc)
        { // 1 если горизонатально
            if (ind && l2 <= wc - w1 && w2 <= lc)
            { // 1
                printf("YES");
                ind = false;
            }
            else if (ind && w2 <= wc - w1 && l2 <= lc)
            { // 2
                printf("YES");
                ind = false;
            }
            else if (ind && l2 <= lc - l1 && w2 <= wc)
            { // 3
                printf("YES");
                ind = false;
            }
            else if (ind && w2 <= lc - l1 && l2 <= wc)
            { // 4
                printf("YES");
                ind = false;
            }
        }
        if (ind && w1 <= lc && l1 <= wc)
        { // 1 если вертикально
            if (ind && l2 <= wc - l1 && w2 <= lc)
            { // 5
                printf("YES");
                ind = false;
            }
            else if (ind && l2 <= lc && w2 <= wc - l1)
            { // 6
                printf("YES");
                ind = false;
            }
            else if (ind && l2 <= lc - w1 && w2 <= wc)
            { // 7
                printf("YES");
                ind = false;
            }
            else if (ind && w2 <= lc - w1 && l2 <= wc)
            { // 8
                printf("YES");
                ind = false;
            }
        }
        if (ind && h1 + h2 <= hc && l1 <= lc && w1 <= wc && l2 <= lc && w2 <= wc)
        {
            printf("YES");
            ind = false;
        }
        else if (ind)
        {
            printf("NO");
        }
    }
    else
    {
        printf("NO");
    }
    return 0;
}
