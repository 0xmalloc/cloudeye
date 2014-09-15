typedef enum _regex_ret_t
{
    REGEX_MATCHED = 0,
    REGEX_MATCHED_NO_VAL = 1,
    REGEX_NOT_MATCHED = 2,
    REGEX_COMPILE_FAILED = 5,
    REGEX_EXEC_FAILED = 6,
    REGEX_PARAM_INVALID = 7,
}regex_ret_t;
regex_ret_t regex_search(const char* pattern, const char* str, char** match_start, unsigned int*  len);
