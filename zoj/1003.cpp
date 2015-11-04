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
        /*��numb�ֽ���ɵ�����£��鿴numa�Ƿ���Էֽ⡣
          �������Ա�֤ĳЩ�����ӱ�numb���ˣ��㲻���ٱ�numa��*/
    }

	if(k == 1 || (f1 && f2)) return;
	if(numa % k == 0) dfs(numa / k, numb, k - 1); //����k,��a�ò���b��
	if(numb % k == 0) dfs(numa, numb / k, k - 1); //����k,��b�ò���a��
	dfs(numa, numb, k - 1); //����k,�Ȳ���a��Ҳ����b��
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


