/*
 * Do the Untwist
 * http://acm.zju.edu.cn/onlinejudge/showProblem.do?problemCode=1006
*/

#include <stdio.h>
#include <string.h>

int main()
{
	int k, ccode[80];
	char ptext[80], ctext[80];
	while (scanf("%d", &k) != EOF&&k)
	{
		scanf("%s", ctext);

		int n = 0;
		while (ctext[n] != '\0')
		{
			if (ctext[n] == '.')
			{
				ccode[n] = 27;
			}
			else if (ctext[n] == '_')
			{
				ccode[n] = 0;
			}
			else
			{
				ccode[n] = ctext[n] - 'a' + 1;
			}
			n++;
		}

		//key
		for (int i = 0; i<n; i++)
		{
			int tem;
			tem = (ccode[i] + i) % 28;
			if (tem == 0)
				ptext[k*i%n] = '_';
			else if (tem == 27)
				ptext[k*i%n] = '.';
			else
				ptext[k*i%n] = tem + 'a' - 1;

		}
		ptext[n] = '\0';

		printf("%s\n", ptext);
	}

	return 0;
}