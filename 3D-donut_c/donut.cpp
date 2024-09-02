#include <cmath>
#include <vector>
#include <iostream>
#include <string>
#include <thread>
#include <chrono>

constexpr int screen_width = 80;
constexpr int screen_height = 22;
constexpr int buffer_size = screen_width * screen_height;
constexpr float pi = 3.14159265358979323846;
constexpr float step_i = 0.02;
constexpr float step_j = 0.07;

void render_frame(float A, float B, std::vector<float>& z, std::vector<char>& b) {
    std::fill(b.begin(), b.end(), ' ');
    std::fill(z.begin(), z.end(), 0);

    for (float j = 0; j < 2 * pi; j += step_j) {
        float cos_j = std::cos(j);
        float sin_j = std::sin(j);

        for (float i = 0; i < 2 * pi; i += step_i) {
            float cos_i = std::cos(i);
            float sin_i = std::sin(i);

            float cos_A = std::cos(A);
            float sin_A = std::sin(A);
            float cos_B = std::cos(B);
            float sin_B = std::sin(B);

            float h = cos_j + 2;
            float D = 1 / (sin_i * h * sin_A + sin_j * cos_A + 5);
            float t = sin_i * h * cos_A - sin_j * sin_A;

            int x = static_cast<int>(screen_width / 2 + screen_width / 2 * D * (cos_i * h * cos_B - t * sin_B));
            int y = static_cast<int>(screen_height / 2 + screen_height / 2 * D * (cos_i * h * sin_B + t * cos_B));
            int o = x + screen_width * y;

            int N = static_cast<int>(8 * ((sin_j * sin_A - sin_i * cos_j * cos_A) * cos_B - sin_i * cos_j * sin_A - sin_j * cos_A - cos_i * cos_j * sin_B));
            if (y < screen_height && y >= 0 && x >= 0 && x < screen_width && D > z[o]) {
                z[o] = D;
                b[o] = "helloworldcisthebest"[N > 0 ? N : 0];
            }
        }
    }
}

void display_frame(const std::vector<char>& b) {
    std::cout << "\x1b[H";
    for (int k = 0; k < buffer_size; ++k) {
        std::cout << (k % screen_width ? b[k] : '\n');
    }
}

int main() {
    float A = 0, B = 0;
    std::vector<float> z(buffer_size);
    std::vector<char> b(buffer_size);
    std::cout << "\x1b[2J";

    while (true) {
        render_frame(A, B, z, b);
        display_frame(b);

        A += 0.04;
        B += 0.02;

        std::this_thread::sleep_for(std::chrono::microseconds(30000));
    }

    return 0;
}
