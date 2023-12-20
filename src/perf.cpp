#include <iostream>
#include <chrono>

int main() {
    // Start the clock
    auto start_time = std::chrono::high_resolution_clock::now();

    // Your loop
    for (int i = 0; i < 10000; ++i) {
        // Do some work in the loop
        std::cout << i*i ;
        std::cout << "\n";
    }

    // Stop the clock
    auto end_time = std::chrono::high_resolution_clock::now();

    // Calculate the duration
    auto duration = std::chrono::duration_cast<std::chrono::microseconds>(end_time - start_time);

    // Output the duration
    std::cout << "Loop execution time: " << duration.count() << " microseconds" << std::endl;

    return 0;
}