#include <iostream>
#include <stdio.h>
#include <cstdlib>
#define size 5

using namespace std;

struct stack{
    int top;
    int items[size];
};

int main(){
    //
    int n, d;
    int opc;
    stack my_stack;

    for(int i=0; i<size;i++){
        cin>>n;
        my_stack.items[i]=n;
    }
    
    return 0;
}

int empty(struct stack *ps){
    if(ps->top==-1){
        return(true);
    }else{
        return(false);
    }
}

int pop(struct stack *ps){
    if(empty (ps)){
        printf("Stack underflow");
        exit(1);
    }
    return (ps->items[ps->top--]);
}

void push(struct stack *ps, int x){
    if(ps->top==size-1){
        printf("Stack overflow");
        exit(1);
    }else{
        ps->items[++(ps->top)]=x;
    }
}