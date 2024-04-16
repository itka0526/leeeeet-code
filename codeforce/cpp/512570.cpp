#include <bits/stdc++.h>

using namespace std;

void solveA() {
    int n;
    cin >> n;
    set<char> seen;
    int points = 0;
    while (n--) {
        char c;
        cin >> c;
        if (seen.find(c) == seen.end()) {
            // +2
            points += 2;
            seen.insert(c);
        } else {
            // +1
            points += 1;
        }
    }
    cout << points << endl;
}

void solveB() {
    int n;
    cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; i++) cin >> a[i];
    set<int> b(a.begin(), a.end());
    int m = b.size();
    int nOps = (n - m + 1) / 2;
    cout << n - nOps * 2 << endl;
}

void solveC() {
    int n;
    cin >> n;
    vector<int> lock(n);
    for (int& it : lock) cin >> it;
    for (int i = 0; i < n; i++) {
        int rotations;
        cin >> rotations;
        while (rotations--) {
            char direction;
            cin >> direction;
            if (direction == 'D') {
                lock[i] += 1;
            } else {
                lock[i] -= 1;
            }
            lock[i] %= 10;
            if (lock[i] < 0) lock[i] += 10;
        }
    }
    for (auto i : lock) std::cout << i << ' ';
    cout << endl;
}

int translateToMins(const string& s) {
    int hh, mm;
    sscanf(s.c_str(), "%d:%d", &hh, &mm);
    return hh * 60 + mm;
}

string formatBack(int mins) {
    char t[6];
    sprintf(t, "%02d:%02d", mins / 60 % 24, mins % 60);
    return string(t);
}

bool validPalindrome(int mins) {
    string s1 = formatBack(mins);
    string s2 = s1;
    reverse(s2.begin(), s2.end());
    return s1 == s2;
}

void solveD() {
    string initialTime;
    int interval;
    cin >> initialTime >> interval;
    int currentTime = translateToMins(initialTime);
    int count = 0;
    // Some loop that updates currentTime on interval and check if that interval has been seen?
    set<string> seen;
    while (seen.find(formatBack(currentTime)) == seen.end()) {
        seen.insert(formatBack(currentTime));
        if (validPalindrome(currentTime)) count++;
        currentTime += interval;
    }
    cout << count << endl;
}

void solveE() {
    int n;
    cin >> n;
    vector<string> parts = {""};
    for (int i = 0; i < n; i++) {
        char ch;
        cin >> ch;
        if (ch == 'W') {
            parts.push_back("");
        } else {
            parts.back() += ch;
        }
    }

    bool ok = true;

    for (const string& s : parts) {
        if (s.size() == 1) {
            ok = false;
            break;
        }
        if (s.size() >= 2) {
            ok &= count(s.begin(), s.end(), 'R') > 0 && count(s.begin(), s.end(), 'B') > 0;
        }
    }
    cout << (ok ? "YES" : "NO") << endl;
}

void solveF() {
    int n;
    cin >> n;
    string s;
    cin >> s;
    int64_t sum = 0;
    vector<int> a(n);
    for (int i = 0; i < n; i++) {
        int left = i, right = n - i - 1;
        int curr = (s[i] == 'L') ? left : right;
        sum += curr;
        // difference between if the facing direction changes and if it did not change
        // For example, say person one is facing left and is standing at 2, and there are 5 people on the right.
        // If the person would change direction he would gain 3 more. 5 - 2 = 3
        // And the difference would be 3 - 1 = 2
        // LLR -> 2 0 2 ->
        //      L:1 = 0, R:1 = 2 |> 2 - 0 -> 2      + 2
        //      L:2 = 1, R:2 = 1 |> 1 - 1 -> 0      Nothing to gain
        //      R:3 = 0, L:3 = 2 |> 2 - 0 -> 2      + 2
        if (s[i] == 'L')
            a[i] = right - left;
        else
            a[i] = left - right;
    }
    sort(a.begin(), a.end(), greater<int>());
    int64_t maxSum = sum;
    for (const int& it : a) {
        sum += it;
        maxSum = max(sum, maxSum);
        cout << maxSum << ' ';
    }
    cout << endl;
}

int main() {
    int t;
    cin >> t;
    while (t--) solveF();
    return 0;
}