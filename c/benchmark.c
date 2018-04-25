/*
 * Determine if two words are an anagram of one another
 * by assigning prime number values to the leters therein,
 * calculating their products, and comparing. If the products
 * are the same, the words are anagrams of one another. 
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

/*
 * unique_value generates a unique value assinged to the 
 * word passed in based on the letters it contains.
 */
int unique_value(const char* word, size_t word_size) {
    int prod = 1;
    for (int i = 0; i < word_size; i++){
        switch (word[i]) {
            case 'a': prod *= 2; break;
            case 'b': prod *= 3; break;
            case 'c': prod *= 5; break;
            case 'd': prod *= 7; break;
            case 'e': prod *= 11; break;
            case 'f': prod *= 13; break;
            case 'g': prod *= 17; break;
            case 'h': prod *= 19; break;
            case 'i': prod *= 23; break;
            case 'j': prod *= 29; break;
            case 'k': prod *= 31; break;
            case 'l': prod *= 37; break;
            case 'm': prod *= 41; break;
            case 'n': prod *= 43; break;
            case 'o': prod *= 47; break;
            case 'p': prod *= 53; break;
            case 'q': prod *= 59; break;
            case 'r': prod *= 61; break;
            case 's': prod *= 67; break;
            case 't': prod *= 71; break;
            case 'u': prod *= 73; break;
            case 'v': prod *= 79; break;
            case 'w': prod *= 83; break;
            case 'x': prod *= 89; break;
            case 'y': prod *= 97; break;
            case 'z': prod *= 101; break;
        }
    }
    return prod;
}

/*
 * is_anagram determines whether the given words are 
 * anagrams of one another. Returns 1 on success and 
 * 0 on failure.
 */
int is_anagram(const char* w1, size_t w1_size, const char* w2, size_t w2_size) {
    if (unique_value(w1, w1_size) == unique_value(w2, w1_size))
        return 1;

    return 0;
}

char *commas(int n){ 
    char *out = malloc(32);
    if (n == 0){
        sprintf(out, "0");
        return out;
    }

    out[31] = 0;
    int i = 31;
    int j = 0;

    while (n != 0){
        i--;
        if (j == 3){
            out[i] = ',';    
            i--;
            j = 0;
        }
        j++;
        out[i] = (n - ((n / 10) * 10)) + '0';
        n /= 10;
    }

    return &out[i];
}

int main() {
    srand(time(NULL)); // seed the random number generator with a time value

    int N = 10000000; // Why is this variable upper case?
    int word_size = 3; // only words that are 3 characters in length

    // create a very long string of random lower case characters.
    int all_size = N*word_size*2; // get the number of sets needs 
    char *str = malloc(all_size); // Why are we doing this?
    for (int i = 0; i < all_size; i++) { // iterate up to the value of all_size 
        str[i] = (rand() % 26) + 'a'; // get the random number, divide by 26, and take the remainder and add 97
    }

    // count the anagrams
    clock_t start = clock();
    int count = 0;
    for (int i = 0; i < all_size; i += word_size * 2) {
        if (is_anagram(&str[i], word_size, &str[i+word_size], word_size)) {
            count++;
        }
    }

    // dur will equal the start substracted from the run of clock now
    // divided by CLOCKS_PER_SEC. Since we're dividing, we need to cast
    // the values to doubles to account for the remainders. The 
    // CLOCKS_PER_SEC on Darwin is defined to be 100.
    double dur = (double)(clock() - start) / (double)CLOCKS_PER_SEC;

    // print out the results of the processing.
    printf("%s matches out of %s compares in %dms (%s ops/sec)\n", 
        commas(count), commas(N), (int)(dur*1000), commas(N/dur));  // wouldn't all of the "comma" calls need to be freed?

    return 0;
}
