#include <iostream>

using namespace std;

int main() {
  int a, b, c;
  cin >> a >> b >> c;

  while (a <= b) {
    if (a % c == 0) {
      cout << a;
      return 0;
    }
    a++;
  }
  cout << "-1";
  return 0;
}