#include <bits/stdc++.h>

using namespace std;

int main() {
    int h, m;
    scanf("%d:%d", &h, &m);
    if (h * 60 + m < 8 * 60 + 30) {
        cout << "MORNING";
    } else if (h * 60 + m < 9 * 60 + 55) {
        cout << 1;
    } else if (h * 60 + m < 10 * 60 + 10) {
        cout << "BREAK";
    } else if (h * 60 + m < 11 * 60 + 35) {
        cout << 2;
    } else if (h * 60 + m < 11 * 60 + 50) {
        cout << "BREAK";
    } else if (h * 60 + m < 13 * 60 + 15) {
        cout << 3;
    } else if (h * 60 + m < 13 * 60 + 45) {
        cout << "BREAK";
    } else if (h * 60 + m < 15 * 60 + 10) {
        cout << 4;
    } else if (h * 60 + m < 15 * 60 + 25) {
        cout << "BREAK";
    } else if (h * 60 + m < 16 * 60 + 50) {
        cout << 5;
    } else if (h * 60 + m < 17 * 60 + 5) {
        cout << "BREAK";
    } else if (h * 60 + m < 18 * 60 + 30) {
        cout << 6;
    } else {
        cout << "EVENING";
    }
    return 0;
}