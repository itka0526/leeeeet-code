/*
    "Самое экстравагантное дупло": геометрия, конструктив, O(n)
    
    Главный подвох: радиусы даны в метрах, а расстояние определяется в сантиметрах. 
    
    Можно пересечь все окружности.
*/

#include <stdio.h>

int main() {
    int n;
    scanf("%d", &n);
    for (double x = 0; n > 0 && x < 1; x += 0.01) {
        for (double y = 0; n > 0 && y < 1; y += 0.01) {
            printf("%.2f %.2f", x, y);
            --n;
        }
    }
    return 0;
}