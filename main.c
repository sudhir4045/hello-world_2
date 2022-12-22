#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

bool binary_search_func(int* arr, int num)
{
    bool value = false;
    int low =0;
    int high = (sizeof(arr))/sizeof(int) - 1;
    int mid = (low+high)/2;
    while(low <= high)
    {
        if(arr[mid] > num)
        {
            high = mid;
            mid = (high + low)/2;
        }
        else if(arr[mid] < num)
        {
            low = mid;
            mid = (high + low)/2;
        }
        else
        {
            value = true;
            break;
        }
    }
    return value;
}

void main(void)
{
    int array_value[] = {1, 34, 45, 456, 567, 678, 876, 980, 1010, 1020};
    bool realval = false;
    realval = binary_search_func(array_value, 456)
    if(realval)
    {
        printf("found value 456\n");
    }

    realval = binary_search_func(array_value, 1)
    if(realval)
    {
        printf("found value 456\n");
    }

    realval = binary_search_func(array_value, 1020)
    if(realval)
    {
        printf("found value 456\n");
    }

    realval = binary_search_func(array_value, 333)
    if(realval)
    {
        printf("found value 456\n");
    }

}
