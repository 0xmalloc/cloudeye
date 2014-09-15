#include<iostream>
#include "util.h"

using namespace std;
int main(int argc, char** argv)
{
    char* match_start = NULL;
    unsigned int match_len = 0;

    const char* logstr = "NOTICE: xx  cost=98  xx";
    const char* pattern = "cost=(\\d+)";

    regex_ret_t ret = regex_search(pattern, logstr, &match_start, &match_len);
    if(ret != REGEX_MATCHED_NO_VAL &&  ret != REGEX_MATCHED){
        cout<< "Not Matched ret:" << ret <<endl;
    }else{
        cout<< "Matched ret:" << ret << endl; 
        cout<< "Matched str:" << string(match_start, match_len) << endl; 
    }
}
