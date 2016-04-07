/*
 *http://acm.zju.edu.cn/onlinejudge/showProblem.do?problemCode=1081
 *2016.03.31
*/
#define _CRT_SECURE_NO_DEPRECATE
#include <stdio.h>
 
struct p{
	int x, y;
}pt[200];
int cmp(int x, int y){
	if (x>y) return 1;
	if (x<y) return -1;
	return 0;
}
int main(){
	int i, j, n, m, cnt, cs = 1;
	int tx, ty, sta, lsta, temp, e;
	//freopen("in.in","r",stdin);  
	while (~scanf("%d", &n) && n){
		scanf("%d", &m);
		for (i = 0; i<n; i++)
			scanf("%d%d", &pt[i].x, &pt[i].y);
		pt[n].x = pt[0].x;
		pt[n].y = pt[0].y;
		if (cs != 1)
			printf("\n");
		printf("Problem %d:\n", cs++);
		for (j = 0; j<m; j++){
			scanf("%d%d", &tx, &ty);
			cnt = 0;
			lsta = cmp(pt[0].y, ty);
			for (i = 1; i <= n; i++){
				e = (pt[i].x - pt[i - 1].x)*(ty - pt[i - 1].y) - (pt[i].y - pt[i - 1].y)*(tx - pt[i - 1].x);
				if (!e && ((tx - pt[i].x)*(tx - pt[i - 1].x) + (ty - pt[i].y)*(ty - pt[i - 1].y) <= 0))  break;
				sta = cmp(pt[i].y, ty);
				temp = sta - lsta;
				if (temp*e<0)
					cnt += temp;
				lsta = sta;
			}
			printf("%s\n", (cnt || i <= n) ? "Within" : "Outside");
		}
	}
	return 0;
}