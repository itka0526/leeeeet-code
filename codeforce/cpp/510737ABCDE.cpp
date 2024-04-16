#include <bits/stdc++.h>

using namespace std;

void solveA() {
    vector<string> board(8);
    for (int i = 0; i < 8; i++) cin >> board[i];
    for (int i = 1; i < 7; i++) {
        for (int j = 1; j < 7; j++) {
            bool ok = board[i][j] == '#' && board[i - 1][j - 1] == '#' && board[i - 1][j + 1] == '#' &&
                      board[i + 1][j - 1] == '#' && board[i + 1][j + 1] == '#';
            if (ok) {
                cout << i + 1 << " " << j + 1 << endl;
                break;
            }
        }
    }
}

void solveB() {
    vector<string> a(8);
    for (int i = 0; i < 8; i++) cin >> a[i];
    bool isRed = false;
    for (int i = 0; i < 8; i++)
        if (a[i] == "RRRRRRRR") isRed = true;
    cout << (isRed ? 'R' : 'B') << endl;
}

void solveCFAILED() {
    int n, m;
    cin >> n >> m;

    vector<long long> diag1(n + m), diag7(n + m);
    vector<vector<long long>> a(n, vector<long long>(m));

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            cin >> a[i][j];
            diag1[i + j] += a[i][j];
            diag7[m + i - j] += a[i][j];
        }
    }

    long long ans = 0LL;

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            ans = max(ans, diag1[i + j] + diag7[m + i - j] - a[i][j]);
        }
    }
    cout << ans << '\n';
}

void solveC() {
    int n, m;
    cin >> n >> m;

    vector<vector<long long>> a(n, vector<long long>(m));
    for (vector<long long> &b : a)
        for (long long &c : b) cin >> c;

    long long ans = 0;

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            long long sum = a[i][j];
            int ci = i, cj = j;

            while (ci >= 0 && cj >= 0) {
                sum += a[ci][cj];
                ci--;
                cj--;
            }

            ci = i, cj = j;
            while (ci < n && cj < m) {
                sum += a[ci][cj];
                ci++;
                cj++;
            }

            ci = i, cj = j;
            while (ci >= 0 && cj < m) {
                sum += a[ci][cj];
                ci--;
                cj++;
            }

            ci = i, cj = j;
            while (ci < n && cj >= 0) {
                sum += a[ci][cj];
                ci++;
                cj--;
            }

            ans = max(ans, sum - 4 * a[i][j]);
        }
    }
    cout << ans << endl;
}

vector<vector<bool>> rotateBy90(vector<vector<bool>> v, int n) {
    vector<vector<bool>> b = v;
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n / 2; j++) {
            swap(b[i][j], b[j][i]);
        }
    }
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n / 2; j++) {
            swap(b[i][j], b[i][n - j - 1]);
        }
    }
    return b;
}

bool areMatricesEqual(vector<vector<bool>> a, vector<vector<bool>> b, int n) {
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            if (a[i][j] != b[i][j]) return false;
        }
    }
    return true;
}

void displayMatrix(vector<vector<bool>> v, string msg = "Matrix : ") {
    cout << msg << endl;

    for (const vector<bool> &it : v) {
        for (bool el : it) {
            cout << el << "  ";
        }
        cout << endl;
    }

    cout << endl;
}

void solveD() {
    int n;
    cin >> n;
    vector<string> a(n);
    for (int i = 0; i < n; i += 1) cin >> a[i];

    int ans = 0;

    for (int row = 0; row < n; row++) {
        for (int col = 0; col < n; col++) {
            int countOnes = 0, countZeros = 0;
            int currRow = row, currCol = col;
            while (1) {
                countOnes += (a[currRow][currCol] == '1');
                countZeros += (a[currRow][currCol] == '0');

                int nextRow = currCol;
                int nextCol = n - currRow - 1;

                currRow = nextRow;
                currCol = nextCol;
                if (currRow == row && currCol == col) break;
            }

            ans += min(countZeros, countOnes);
        }
    }
    cout << ans / 4 << endl;
}

void solveE() {
    vector<vector<long long>> s(1024, vector(1024, 0LL));

    int n, q;
    cin >> n >> q;

    while (n--) {
        int h, w;
        cin >> h >> w;
        s[h][w] += h * w;
    }

    for (int r = 1; r < 1024; r++) {
        for (int c = 1; c < 1024; c++) {
            s[r][c] += s[r - 1][c] + s[r][c - 1] - s[r - 1][c - 1];
        }
    }

    while (q--) {
        int h1, w1, h2, w2;
        cin >> h1 >> w1 >> h2 >> w2;
        h2--, w2--;
        cout << s[h2][w2] - s[h1][w2] - s[h2][w1] + s[h1][w1] << '\n';
    }
}

int main() {
    int q;
    cin >> q;
    while (q--) solveE();
    return 0;
}