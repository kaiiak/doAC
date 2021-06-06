//http://acm.zju.edu.cn/onlinejudge/showProblem.do?problemCode=1003

#include<cstdio>
#include<algorithm>
using namespace std;

bool f1, f2;

void dfs(int numa, int numb, int k)
{
	if(numb == 1)
    {
        f2 = true;
        if(numa == 1) f1 = true;
        /*在numb分解完成的情况下，查看numa是否可以分解。
          这样可以保证某些公因子被numb用了，便不能再被numa用*/
    }

	if(k == 1 || (f1 && f2)) return;
	if(numa % k == 0) dfs(numa / k, numb, k - 1); //因子k,被a用不被b用
	if(numb % k == 0) dfs(numa, numb / k, k - 1); //因子k,被b用不被a用
	dfs(numa, numb, k - 1); //因子k,既不被a用也不被b用
}

int main()
{
	int a, b;
	while(scanf("%d%d",&a,&b) != EOF)
	{
		if(a < b) swap(a, b);
		f1 = f2 = false;
		dfs(a, b, 100);
		if(!f1 && f2) printf("%d\n",b);
		else printf("%d\n",a);
	}
	return 0;
}


