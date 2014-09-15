#include <pcre.h>
#include <string>
#include "util.h"

#define OVECCOUNT 30    /* should be a multiple of 3 */

regex_ret_t regex_search(const char* pattern, const char* str, char** match_start, unsigned int*  len)
{
    if(NULL == str)
        return REGEX_PARAM_INVALID;

    //init
    *len = 0;
    pcre *re = NULL;
    const char* error = NULL;
    int erroffset;
    regex_ret_t ret;
    int rc,i;
    int ovector[OVECCOUNT];
    //re = pcre_compile(pattern, PCRE_UNGREEDY, &error, &erroffset, NULL); // PCRE_UNGREEDY  非贪婪模式
    re = pcre_compile(pattern, 0, &error, &erroffset, NULL); //贪婪模式在有些情况会出问题，所以如果需要使用,请在正则中指定 “?U”
    if (re == NULL) {
        //DSTREAM_WARN( "PCRE compilation failed at offset %d: %s\n", erroffset, error);
        return REGEX_COMPILE_FAILED;;
    }

    //match
    rc = pcre_exec(re, NULL, str, strlen(str), 0, 0, ovector, OVECCOUNT);
    if (rc < 0) {
        if (rc == PCRE_ERROR_NOMATCH){
            ret = REGEX_NOT_MATCHED;;
            goto FREE_RETURN;
        }
        else{
            //DSTREAM_WARN( "PCRE pcre_exec failed, matching error %d\n", rc);
            ret = REGEX_EXEC_FAILED;
            goto FREE_RETURN;
        }
    }

    //get result
    if(2 == rc){
        i = 1; //直取第二个, 有点trick
        *match_start = const_cast<char*>(str + ovector[2*i]);
        *len = ovector[2*i+1] - ovector[2*i];
        ret = REGEX_MATCHED;
    }else{
        ret = REGEX_MATCHED_NO_VAL;
    }

FREE_RETURN:
    pcre_free(re);
    return ret;
}
