#include<iostream>
#include "util.h"

using namespace std;
int main(int argc, char** argv)
{
    char* match_start = NULL;
    unsigned int match_len = 0;

    const char* logstr = "NOTICE: rpc_service  cost=98 ispv=1 method=add conn_time=11 read_time=5 write_time=19";
    const char* pattern[] = {
        "method=(\\w+)",
        "cost=(\\d+)",
        "rpc_service",
        "error",
        "rpc_time=(\\d+)",
        "rpc_method=(\\w+)",
    };
    for(int i = 0; i < 6; i++){
        cout<<"pattern:" << pattern[i] << endl;
        regex_ret_t ret = regex_search(pattern[i], logstr, &match_start, &match_len);
        if(ret != REGEX_MATCHED_NO_VAL &&  ret != REGEX_MATCHED){
            cout<< "Not Matched ret:" << ret <<endl;
        }else{
            cout<< "Matched ret:" << ret << endl; 
            cout<< "Matched str:" << string(match_start, match_len) << endl; 
        }   
    }   
}
