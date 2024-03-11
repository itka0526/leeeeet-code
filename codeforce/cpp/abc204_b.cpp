#include <iostream>

using namespace std;

int main() {
  int n;
  cin >> n;
  int a[n];
  int c = 0;
  int total = 0;
  while (c < n) {
    int t;
    cin >> t;
    if (t > 10) {
      total += t - 10;
      a[c] = 10;
    } else {
      a[c] = t;
    }
    c++;
  };
  cout << total << "\n";
}