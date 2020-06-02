#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <math.h>

void reverse(char text[]);
void clean(char text[]);

int main(void) {
    int lines_number = 0;
    int iter = 0;

    scanf("%d", &lines_number);

    int ch;
    while((ch = getchar()) == 32);

    char text[lines_number][1000];

    while (iter < lines_number) {
        fgets(text[iter], 1000, stdin);

        // Removes \n
        clean(text[iter]);

        if (text[iter] == NULL) {
            fprintf(stderr, "Error reading input.\n");
            return -1;
        }

        size_t i;

        for (i = 0; i < strlen(text[iter]); i++) {
            if (isalpha(text[iter][i])) {
            if (!isspace(text[iter][i]))
                text[iter][i] = text[iter][i]+3;
            }
        }

        reverse(text[iter]);

        for (i = floor(strlen(text[iter])/2); i < strlen(text[iter]); i++) {
            text[iter][i] = text[iter][i]-1;
        }

        printf("%s\n", text[iter]);

        iter++;
    }

    return 0;
}

void reverse(char text[]) {
    size_t i;
    size_t j;

    for (i = 0, j = strlen(text)-1; i < j; i++, j--) {
        char c = text[i];
        text[i] = text[j];
        text[j] = c;
    }
}

void clean(char text[]) {
    size_t i;

    for (i = 0; i < strlen(text); i++) {
        if (text[i] == '\n') {
            text[i] = '\000';
        }
    }
}
