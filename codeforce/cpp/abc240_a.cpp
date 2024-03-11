#include <iostream>

using namespace std;

int main() {
  // cout << "hello world!"
  //      << "\n";

  int a, b;
  cin >> a >> b;

  // cout << a << b;
  if ((a == 1 && b == 10) || a + 1 == b) {
    cout << "Yes";
  } else {
    cout << "No";
  }
}