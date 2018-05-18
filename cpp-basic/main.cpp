#include <iostream>
#include <fstream>
#include <vector>

int
main(int argc, char** argv)
{
    if (argc < 2)
        return 1;

    std::ifstream in(argv[1], std::ifstream::binary);
    if (!in)
        return 2;

    in.seekg(0, in.end);
    std::size_t n = in.tellg();
    in.seekg(0, in.beg);

    std::vector<int> buf(n / sizeof(int));
    in.read((char*)&buf[0], n);
    in.close();

    int acc = 0;
    for (int i = 0; i < n/sizeof(int); i++)
        if (buf[i] == -1)
            acc++;

    std::cout << acc << "/" << n/4 << std::endl;
    return 0;
}
