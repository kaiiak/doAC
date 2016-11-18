/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *twoSum(int *nums, int numsSize, int target)
{
    for (int i = 0; i < numsSize; i++)
    {
	int temp = target - nums[i];
	for (int j = i + 1; j < numsSize; j++)
	{
	    if (nums[j] == temp)
	    {
		int *b = (int *)malloc(2 * sizeof(int));
		b[0] = i;
		b[1] = j;
		printf("%d,%d", i, j);
		return b;
	    }
	}
    }
    return;
}
