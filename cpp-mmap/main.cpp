#include <iostream>
#include <fstream>
#include <vector>
#include <sys/mman.h>

#define MB (1 * 1024 * 1024)

int
main(int argc, char** argv)
{
    if (argc < 2)
        return 1;

    FILE* in = fopen(argv[1], "rb");
    int* buf = (int*)mmap(0, 256 * MB, PROT_READ, MAP_FILE | MAP_PRIVATE,
                          fileno(in), 0);

    int n = (256 * MB) / sizeof(int);
    int acc = 0;
    for (int i = 0; i < n; i++)
        if (buf[i] == -1)
            acc++;

    std::cout << acc << "/" << n << std::endl;
    return 0;
}
