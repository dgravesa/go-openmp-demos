#
#include <iostream>
#include <string>
#include <vector>
#include <chrono>
#include <cstdlib>
#include <cmath>
#include <omp.h>
#include "../include/Opts.hpp"

const Option options[] = {
    {"help", 'h', NO_ARG, "print this usage dialog"},
    {"num-items", 'N', REQUIRED_ARG, "number of array items (default: 1000000)"},
    {"num-threads", 't', REQUIRED_ARG, "number of OpenMP threads to use (default: max)"},
    {0, 0, 0, 0}
};

double sinc(double x) {
    if (x == 0.0) {
        return 1.0;
    }
    return sin(x) / x;
}

int main(int argc, char** argv) {
    OptHandler optHandler(options, argc, argv);
    int N = 1000000, numThreads = omp_get_max_threads();

    // read command line arguments
    char optCode;
    std::string optArg;
    while (optHandler.getOpt(optCode, optArg)) {
        switch (optCode) {
        case 'N':
            N = std::atoi(optArg.c_str());
            break;
        case 't':
            numThreads = std::atoi(optArg.c_str());
            break;
        case 'h':
        default:
            optHandler.printUsage("usage: ./0-cpp-demo");
            return 1;
        }
    }

    // allocate vectors and generate random input
    std::vector<double> inputVec(N), outputVec(N);
    for (int i = 0; i < N; i++) {
        double randNormVal = rand() / (double) RAND_MAX;
        inputVec[i] = 10.0 * (randNormVal - 0.5); // -5 to 5
    }

    // execute loop
    omp_set_num_threads(numThreads);
    auto startTime = std::chrono::system_clock::now();
    #pragma omp parallel for
    for (int i = 0; i < N; i++) {
        outputVec[i] = sinc(inputVec[i] * M_PI);
    }
    auto stopTime = std::chrono::system_clock::now();

    // print output time
    float seconds = std::chrono::duration<float>(stopTime - startTime).count();
    if (seconds < 1.0f) {
        std::cout << 1000 * seconds << "ms" << std::endl;
    } else {
        std::cout << seconds << "s" << std::endl;
    }

    return 0;
}
