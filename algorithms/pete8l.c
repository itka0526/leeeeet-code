
int flag[2];
int turn;

void p0() {
    while(1) {
        flag[0] = 1;
        turn = 1;
        while(flag[1] && turn == 1) {
            flag[0] = 0;
        }
    }
}

void p1() {
    while(1) {
        flag[1] = 1;
        turn = 0;
        while(flag[0] && turn == 0) {
            flag[1] = 0;
        }
    }
}

int main() {
    flag[0] = 0;
    flag[1] = 0;
    parbegin(p0, p1);
    return 0;
}