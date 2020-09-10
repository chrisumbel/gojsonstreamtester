#include <iostream>

int main() {
    int i;
    using namespace std;

    while(1) {
        std::cout << "{\"Value\": " << i++ << ", \"Pom\": {\"OtherThing\": \"OtherValue\"}}" << endl;
    }

    return 0;
}