#

# NOTE: using this because clang OpenMP isn't working on my Mac
CXX := g++-10
CXXFLAGS := -O3 -std=c++11 -fopenmp

CXX_DEMOS := $(addprefix bin/,$(shell ls demos | grep "cpp-"))

all: go cpp

bin:
	mkdir -p bin

go: bin
	go build -o bin ./...

cpp: $(CXX_DEMOS)

bin/cpp-%: demos/cpp-%/main.cpp bin
	$(CXX) $(CXXFLAGS) $< -o $@

clean:
	rm -rf bin

