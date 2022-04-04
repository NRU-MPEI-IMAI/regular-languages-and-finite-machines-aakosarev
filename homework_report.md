







## Теоретические модели вычислений

***

## ДЗ  1

## Косарев А.А.

## А-05-19

___

## Задача 1 : Построить конечный автомат, распознающий язык. Автомат должен быть детерминированным.

Язык № 1:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b,c%20%5C%7D%5E*%20%7C%20%7Cw%7C_c%20=%201%20%5C%7D)

Автомат № 1:

![](./first_task/svg_files/gr_1_1.svg)

Язык № 2:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5Cleq%202;%20%7Cw%7C_b%20%5Cgeq%202%5C%7D)

Автомат № 2:

![](./first_task/svg_files/gr_1_2.svg)

Язык № 3:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5Cneq%20%7Cw%7C_b%20%5C%7D)

Автомат № 3:

Автомат функционирует только если a и b чередуются. В противном случае, предполагаю, автомат построить невозможно.

![](./first_task/svg_files/gr_1_3.svg)

Язык № 4:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20ww%20=%20www%20%5C%7D)

Автомат № 4:

![](./first_task/svg_files/gr_1_4.svg)

***

## Задача 2 : Построить конечный автомат, используя прямое произведение.

Язык № 1:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%202%20%5Cwedge%20%7Cw%7C_b%20%5Cgeq%202%5C%7D)

Разобьём на два автомата:

Автомат № 1:

![](./second_task/svg_files/gr_2_1_1.svg)

Автомат № 2:

![](./second_task/svg_files/gr_2_1_2.svg)

![](https://latex.codecogs.com/svg.image?%5CSigma_1%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_1%20=%20%5C%7Bq_1,q_2,q_3%5C%7D%5C%5Cs_1=%5C%7Bq_1%5C%7D%5C%5CT_1=%5C%7Bq_3%5C%7D%5C%5C%5Cdelta_1=%20)

| -   | a   | b   |
| --- | --- | --- |
| q_1 | q_2 | q_1 |
| q_2 | q_3 | q_2 |
| q_3 | q_3 | q_3 |

![](https://latex.codecogs.com/svg.image?%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_2%20=%20%5C%7Bp_1,p_2,p_3%5C%7D%5C%5Cs_2%20=%20%5C%7Bp_1%5C%7D%5C%5CT_2%20=%20%5C%7Bp_3%5C%7D%5C%5C%5Cdelta_2%20=%20%20%20)

| -   | a   | b   |
| --- | --- | --- |
| p_1 | p_1 | p_2 |
| p_2 | p_2 | p_3 |
| p_3 | p_3 | p_3 |

![](https://latex.codecogs.com/svg.image?%5CSigma%20=%20%5CSigma_1%20%5Ccup%20%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5C%5C%5CQ=%20Q_1%5Ctimes%20Q_2=%5C%7B%3Cq_1,p_1%3E,%3Cq_1,p_2%3E,%3Cq_1,p_3%3E,%3Cq_2,p_1%3E,%3Cq_2,p_2%3E,%3Cq_2,p_3%3E,%3Cq_3,p_1%3E,%3Cq_3,p_2%3E,%3Cq_3,p_3%3E%5C%7D%5C%5C%5C%5Cs%20=%20%3Cs_1,s_2%3E%20=%20%3Cq_1,p_1%3E%5C%5C%5C%5CT=T_1%5Ctimes%20T_2=%3Cq_3,p_3%3E%5C%5C%5C%5C%5Cdelta=%20)

| -         | a         | b         |
|:--------- | --------- | --------- |
| <q_1,p_1> | <q_2,p_1> | <q_1,p_2> |
| <q_1,p_2> | <q_2,p_2> | <q_1,p_3> |
| <q_1,p_3> | <q_2,p_3> | <q_1,p_3> |
| <q_2,p_1> | <q_3,p_1> | <q_2,p_2> |
| <q_2,p_2> | <q_3,p_2> | <q_2,p_3> |
| <q_2,p_3> | <q_3,p_3> | <q_2,p_3> |
| <q_3,p_1> | <q_3,p_1> | <q_3,p_2> |
| <q_3,p_2> | <q_3,p_2> | <q_3,p_3> |
| <q_3,p_3> | <q_3,p_3> | <q_3,p_3> |

Результирующий автомат, полученный прямым произведением вышеизложенных автоматов:

![](./second_task/svg_files/gr_2_1.svg)

Язык № 2:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%203%20%5Cwedge%20%7Cw%7C_b%20%5Ctext%7B%20is%20odd%7D%20%5C%7D%20)

Разобьём на два автомата:

Автомат № 1:

![](./second_task/svg_files/gr_2_2_1.svg)

Автомат № 2:

![](./second_task/svg_files/gr_2_2_1.svg)

![](https://latex.codecogs.com/svg.image?%5CSigma_1%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_1%20=%20%5C%7Bq_1,q_2,q_3,q_4%5C%7D%5C%5Cs_1=%5C%7Bq_1%5C%7D%5C%5CT_1=%5C%7Bq_4%5C%7D%5C%5C%5Cdelta_1=%20)

| -   | a   | b   |
| --- | --- | --- |
| q_1 | q_2 | q_1 |
| q_2 | q_3 | q_2 |
| q_3 | q_4 | q_3 |
| q_4 | q_4 | q_4 |

![](https://latex.codecogs.com/svg.image?%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_2%20=%20%5C%7Bp_1,p_2%5C%7D%5C%5Cs_2=%5C%7Bp_1%5C%7D%5C%5CT_2=%5C%7Bp_2%5C%7D%5C%5C%5Cdelta_2=%20)

| -   | a   | b   |
| --- | --- | --- |
| p_1 | p_1 | p_2 |
| p_2 | p_2 | p_1 |

![](https://latex.codecogs.com/svg.image?%5CSigma%20=%20%5CSigma_1%20%5Ccup%20%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5C%5C%5CQ=%20Q_1%5Ctimes%20Q_2=%5C%7B%3Cq_1,p_1%3E,%3Cq_1,p_2%3E,%3Cq_2,p_1%3E,%3Cq_2,p_2%3E,%3Cq_3,p_1%3E,%3Cq_3,p_2%3E,%3Cq_4,p_1%3E,%3Cq_4,p_2%3E%5C%7D%5C%5C%5C%5Cs%20=%20%3Cs_1,s_2%3E%20=%20%3Cq_1,p_1%3E%5C%5C%5C%5CT=T_1%5Ctimes%20T_2=%3Cq_4,p_2%3E%5C%5C%5C%5C%5Cdelta=)

| -         | a         | b         |
| --------- | --------- | --------- |
| <q_1,p_1> | <q_2,p_1> | <q_1,p_2> |
| <q_1,p_2> | <q_2,p_2> | <q_1,p_1> |
| <q_2,p_1> | <q_3,p_1> | <q_2,p_2> |
| <q_2,p_2> | <q_3,p_2> | <q_2,p_1> |
| <q_3,p_1> | <q_4,p_1> | <q_3,p_2> |
| <q_3,p_2> | <q_4,p_2> | <q_3,p_1> |
| <q_4,p_1> | <q_4,p_1> | <q_4,p_2> |
| <q_4,p_2> | <q_4,p_2> | <q_4,p_1> |

Результирующий автомат, полученный прямым произведением вышеизложенных автоматов:

![](./second_task/svg_files/gr_2_2.svg)

Язык № 3:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5Ctext%7B%20is%20even%7D%20%5Cwedge%20%7Cw%7C_b%20%5Ctext%7B%20multiple%20of%203%7D%5C%7D)

Разобьём на два автомата:

Автомат № 1:

![](./second_task/svg_files/gr_2_3_1.svg)

Автомат № 2:

![](./second_task/svg_files/gr_2_3_2.svg)

![](https://latex.codecogs.com/svg.image?%5CSigma_1%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_1%20=%20%5C%7Bq_1,q_2%5C%7D%5C%5Cs_1=%5C%7Bq_1%5C%7D%5C%5CT_1=%5C%7Bq_1%5C%7D%5C%5C%5Cdelta_1=)

| -   | a   | b   |
| --- | --- | --- |
| q_1 | q_2 | q_1 |
| q_2 | q_1 | q_2 |

![](https://latex.codecogs.com/svg.image?%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_2%20=%20%5C%7Bp_1,p_2,p_3%5C%7D%5C%5Cs_2=%5C%7Bp_1%5C%7D%5C%5CT_2=%5C%7Bp_1%5C%7D%5C%5C%5Cdelta_2=)

| -   | a   | b   |
| --- | --- | --- |
| p_1 | p_1 | p_2 |
| p_2 | p_2 | p_3 |
| p_3 | p_3 | p_1 |

![](https://latex.codecogs.com/svg.image?%5CSigma%20=%20%5CSigma_1%20%5Ccup%20%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5C%5C%5CQ=%20Q_1%5Ctimes%20Q_2=%5C%7B%3Cq_1,p_1%3E,%3Cq_1,p_2%3E,%3Cq_1,p_3%3E,%3Cq_2,p_1%3E,%3Cq_2,p_2%3E,%3Cq_2,p_3%3E%5C%7D%5C%5C%5C%5Cs%20=%20%3Cs_1,s_2%3E%20=%20%3Cq_1,p_1%3E%5C%5C%5C%5CT=T_1%5Ctimes%20T_2=%3Cq_1,p_1%3E%5C%5C%5C%5C%5Cdelta=)

| -         | a         | b         |
| --------- | --------- | --------- |
| <q_1,p_1> | <q_2,p_1> | <q_1,p_2> |
| <q_1,p_2> | <q_2,p_2> | <q_1,p_3> |
| <q_1,p_3> | <q_2,p_3> | <q_1,p_1> |
| <q_2,p_1> | <q_1,p_1> | <q_2,p_2> |
| <q_2,p_2> | <q_1,p_2> | <q_2,p_3> |
| <q_2,p_3> | <q_1,p_3> | <q_2,p_1> |

Результирующий автомат, полученный прямым произведением вышеизложенных автоматов:

![](./second_task/svg_files/gr_2_3.svg)

Язык № 4:

![](https://latex.codecogs.com/svg.image?L%20=%20%5Coverline%7BL_3%7D)

Заменим терминальные состояния на терминальные и наоборот и получим требуемый автомат:

![](./second_task/svg_files/gr_2_4.svg)

Язык № 5:

![](https://latex.codecogs.com/svg.image?L%20=%20L_2%20/%20L_3)

Раскроем разность через пересечение:

![](https://latex.codecogs.com/svg.image?L%20=%20L_2%20%5Ccap%20%5Coverline%7BL_3%7D)

Найдем прямое произведение

Автомат № 1:

![](https://latex.codecogs.com/svg.image?%5CSigma_1%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_1%20=%20%5C%7B%3Cq_1,p_1%3E,%3Cq_1,p_2%3E,%3Cq_2,p_1%3E,%3Cq_2,p_2%3E,%3Cq_3,p_1%3E,%3Cq_3,p_2%3E,%3Cq_4,p_1%3E,%3Cq_4,p_2%3E%5C%7D%5C%5Cs_1=%5C%7B%3Cq_1,p_1%3E%5C%7D%5C%5CT_1=%5C%7B%3Cq_4,p_2%3E%5C%7D%5C%5C%5Cdelta_1=)

(в функции перехода сразу переименуем состояния, чтобы в последующем не писать по 4 буквы)

| -   | a   | b   |
| --- | --- | --- |
| f_1 | f_3 | f_2 |
| f_2 | f_4 | f_1 |
| f_3 | f_5 | f_4 |
| f_4 | f_6 | f_3 |
| f_5 | f_7 | f_6 |
| f_6 | f_8 | f_5 |
| f_7 | f_7 | f_8 |
| f_8 | f_8 | f_7 |

![](./second_task/svg_files/gr_2_5_1.svg)

Автомат № 2:

![](https://latex.codecogs.com/svg.image?%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5CQ_2%20=%20%5C%7B%3Cz_1,r_1%3E,%3Cz_1,r_2%3E,%3Cz_1,r_3%3E,%3Cz_2,r_1%3E,%3Cz_2,r_2%3E,%3Cz_2,r_3%3E%5C%7D%5C%5Cs_2=%5C%7B%3Cz_1,r_1%3E%5C%7D%5C%5CT_2=%5C%7B%3Cz_1,r_2%3E,%3Cz_1,r_3%3E,%3Cz_2,r_1%3E,%3Cz_2,r_2%3E,%3Cz_2,r_3%3E%5C%7D%5C%5C%5Cdelta_2=)

(в функции перехода сразу переименуем состояния, чтобы в последующем не писать по 4 буквы)

| -   | a   | b   |
| --- | --- | --- |
| g_1 | g_4 | g_2 |
| g_2 | g_5 | g_3 |
| g_3 | g_6 | g_1 |
| g_4 | g_1 | g_5 |
| g_5 | g_2 | g_6 |
| g_6 | g_3 | g_4 |

![](./second_task/svg_files/gr_2_5_2.svg)

![](https://latex.codecogs.com/svg.image?%5CSigma%20=%20%5CSigma_1%20%5Ccup%20%5CSigma_2%20=%20%5C%7Ba,b%5C%7D%5C%5C%5C%5CQ=%20Q_1%5Ctimes%20Q_2=%5C%7B%3Cf_1,g_1%3E,%3Cf_1,g_2%3E,%3Cf_1,g_3%3E,%3Cf_1,g_4%3E,%3Cf_1,g_5%3E,%3Cf_1,g_6%3E,%3Cf_2,g_1%3E,%3Cf_2,g_2%3E,%3Cf_2,g_3%3E,%3Cf_2,g_4%3E,%3Cf_2,g_5%3E,%3Cf_2,g_6%3E,%3Cf_3,g_1%3E,%3Cf_3,g_2%3E,%3Cf_3,g_3%3E,%3Cf_3,g_4%3E,%3Cf_3,g_5%3E,%3Cf_3,g_6%3E,%3Cf_4,g_1%3E,%3Cf_4,g_2%3E,%3Cf_4,g_3%3E,%3Cf_4,g_4%3E,%3Cf_4,g_5%3E,%3Cf_4,g_6%3E,%3Cf_5,g_1%3E,%3Cf_5,g_2%3E,%3Cf_5,g_3%3E,%3Cf_5,g_4%3E,%3Cf_5,g_5%3E,%3Cf_5,g_6%3E,%3Cf_6,g_1%3E,%3Cf_6,g_2%3E,%3Cf_6,g_3%3E,%3Cf_6,g_4%3E,%3Cf_6,g_5%3E,%3Cf_6,g_6%3E,%3Cf_7,g_1%3E,%3Cf_7,g_2%3E,%3Cf_7,g_3%3E,%3Cf_7,g_4%3E,%3Cf_7,g_5%3E,%3Cf_7,g_6%3E,%3Cf_8,g_1%3E,%3Cf_8,g_2%3E,%3Cf_8,g_3%3E,%3Cf_8,g_4%3E,%3Cf_8,g_5%3E,%3Cf_8,g_6%3E%5C%7D%5C%5C%5C%5Cs%20=%20%3Cs_1,s_2%3E%20=%20%3Cf_1,g_1%3E%5C%5C%5C%5CT=T_1%5Ctimes%20T_2=%5C%7B%3Cf_8,g_2%3E,%3Cf_8,g_3%3E,%3Cf_8,g_4%3E,%3Cf_8,g_5%3E,%3Cf_8,g_6%3E%5C%7D%5C%5C%5C%5C%5Cdelta=)

| -         | a         | b         |
| --------- | --------- | --------- |
| <f_1,g_1> | <f_3,g_4> | <f_2,g_2> |
| <f_1,g_2> | <f_3,g_5> | <f_2,g_3> |
| <f_1,g_3> | <f_3,g_6> | <f_2,g_1> |
| <f_1,g_4> | <f_3,g_1> | <f_2,g_5> |
| <f_1,g_5> | <f_3,g_2> | <f_2,g_6> |
| <f_1,g_6> | <f_3,g_3> | <f_2,g_4> |
| <f_2,g_1> | <f_4,g_4> | <f_1,g_2> |
| <f_2,g_2> | <f_4,g_5> | <f_1,g_3> |
| <f_2,g_3> | <f_4,g_6> | <f_1,g_1> |
| <f_2,g_4> | <f_4,g_1> | <f_1,g_5> |
| <f_2,g_5> | <f_4,g_2> | <f_1,g_6> |
| <f_2,g_6> | <f_4,g_3> | <f_1,g_4> |
| <f_3,g_1> | <f_5,g_4> | <f_4,g_2> |
| <f_3,g_2> | <f_5,g_5> | <f_4,g_3> |
| <f_3,g_3> | <f_5,g_6> | <f_4,g_1> |
| <f_3,g_4> | <f_5,g_1> | <f_4,g_5> |
| <f_3,g_5> | <f_5,g_2> | <f_4,g_6> |
| <f_3,g_6> | <f_5,g_3> | <f_4,g_5> |
| <f_4,g_1> | <f_6,g_4> | <f_3,g_2> |
| <f_4,g_2> | <f_6,g_5> | <f_3,g_3> |
| <f_4,g_3> | <f_6,g_6> | <f_3,g_1> |
| <f_4,g_4> | <f_6,g_1> | <f_3,g_5> |
| <f_4,g_5> | <f_6,g_2> | <f_3,g_6> |
| <f_4,g_6> | <f_6,g_3> | <f_3,g_4> |
| <f_5,g_1> | <f_7,g_4> | <f_6,g_2> |
| <f_5,g_2> | <f_7,g_5> | <f_6,g_3> |
| <f_5,g_3> | <f_7,g_6> | <f_6,g_1> |
| <f_5,g_4> | <f_7,g_1> | <f_6,g_5> |
| <f_5,g_5> | <f_7,g_2> | <f_6,g_6> |
| <f_5,g_6> | <f_7,g_3> | <f_6,g_4> |
| <f_6,g_1> | <f_8,g_4> | <f_5,g_2> |
| <f_6,g_2> | <f_8,g_5> | <f_5,g_3> |
| <f_6,g_3> | <f_8,g_6> | <f_5,g_1> |
| <f_6,g_4> | <f_8,g_1> | <f_5,g_5> |
| <f_6,g_5> | <f_8,g_2> | <f_5,g_6> |
| <f_6,g_6> | <f_8,g_3> | <f_5,g_4> |
| <f_7,g_1> | <f_7,g_4> | <f_8,g_2> |
| <f_7,g_2> | <f_7,g_5> | <f_8,g_3> |
| <f_7,g_3> | <f_7,g_6> | <f_8,g_1> |
| <f_7,g_4> | <f_7,g_1> | <f_8,g_5> |
| <f_7,g_5> | <f_7,g_2> | <f_8,g_6> |
| <f_7,g_6> | <f_7,g_3> | <f_8,g_4> |
| <f_8,g_1> | <f_8,g_4> | <f_7,g_2> |
| <f_8,g_2> | <f_8,g_5> | <f_7,g_3> |
| <f_8,g_3> | <f_8,g_6> | <f_7,g_1> |
| <f_8,g_4> | <f_8,g_1> | <f_7,g_5> |
| <f_8,g_5> | <f_8,g_2> | <f_7,g_6> |
| <f_8,g_6> | <f_8,g_3> | <f_7,g_4> |

Результирующий автомат, полученный прямым произведением вышеизложенных автоматов:

![](./second_task/svg_files/gr_2_5.svg)

## Задача 3 : Построить минимальный ДКА по регулярному выражению.

Регулярное выражение № 1:

![](https://latex.codecogs.com/svg.image?(ab&plus;aba)%5E*a)

Сперва построим недетерминированный конечный автомат, а затем преобразуем его в детерминированный.

НКА:

![](./third_task/svg_files/gr_3_1_1.svg)

Теперь составим таблицу для преобразования НКА в ДКА:

|              | a            | b     |
| ------------ | ------------ | ----- |
| q1           | q3,q7,q10    | -     |
| q3,q7,q10    | -            | q4,q8 |
| q4,q8        | q3,q5,q10,q7 | -     |
| q3,q5,q10,q7 | q3,q7,q10    | q4,q8 |

Получаем ДКА:

![](./third_task/svg_files/gr_3_1.svg)

Регулярное выражение № 2:

![](https://latex.codecogs.com/svg.image?a(a(ab)%5E*b)%5E*(ab)%5E*)

Сперва построим недетерминированный конечный автомат, а затем преобразуем его в детерминированный.

НКА:

![](./third_task/svg_files/gr_3_2_1.svg)

Теперь составим таблицу для преобразования НКА в ДКА:

|       | a     | b     |
| ----- | ----- | ----- |
| q1    | q2    | -     |
| q2    | q3    | -     |
| q3    | q4    | q6    |
| q4    | -     | q5    |
| q6    | q3,q7 | -     |
| q5    | q4    | q6    |
| q3,q7 | q4    | q6,q8 |
| q6,q8 | q3,q7 | -     |



Получаем ДКА:

![](./third_task/svg_files/gr_3_2.svg)

Регулярное выражение № 3:

![](https://latex.codecogs.com/svg.image?(a&plus;(a&plus;b)(a&plus;b)b)%5E*)

Сперва построим недетерминированный конечный автомат, а затем преобразуем его в детерминированный.

НКА:

![](./third_task/svg_files/gr_3_3_1.svg)

Теперь составим таблицу для преобразования НКА в ДКА:

|          | a        | b        |
| -------- | -------- | -------- |
| q1       | q1,q2    | q2       |
| q1,q2    | q1,q2,q3 | q2,q3    |
| q2       | q3       | q3       |
| q1,q2,q3 | q1,q2,q3 | q1,q2,q3 |
| q2,q3    | q3       | q1,q3    |
| q3       | -        | q1       |
| q1,q3    | q1,q2    | q1,q2    |

Получаем ДКА:

![](./third_task/svg_files/gr_3_3.svg)

Регулярное выражение № 4:

![](https://latex.codecogs.com/svg.image?(b&plus;c)((ab)%5E*c&plus;(ba)%5E*)%5E*)

Можем сразу построить ДКА:

![](./third_task/svg_files/gr_3_4.svg)

## Задача 4 : Определить является ли язык регулярным или нет

Язык № 1:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7B(aab)%5Enb(aba)%5Em%20%5Cmid%20n%20%5Cgeq%20%200,%20m%20%5Cgeq%20%200%5C%7D)

![](./fourth_task/svg_files/gr_4_1.svg)

Язык является регулярным.



Язык № 2:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Buaav%20%5Cmid%20u%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20v%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20%7Cu%7C_b%20%5Cgeqslant%20%7Cv%7C_a%5C%7D)

Лемма о накачке:

![](https://latex.codecogs.com/svg.image?L-%5Ctext%7Bregular%7D%5CRightarrow%20%5Cexists%20n%5C;%5C;%5C;%5Cforall%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cexists%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cforall%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cin%20L)

Отрицание:

![](https://latex.codecogs.com/svg.image?%5Cforall%20n%5C;%5C;%5C;%5Cexists%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cforall%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cexists%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cnotin%20%20L)

Возьмём слово

![](https://latex.codecogs.com/svg.image?%5Comega%20=%20b%5Enaaa%5En,%20%5Cforall%20n%5C%5C%5Cleft%7C%5Comega%20%20%5Cright%7C=2n&plus;2%20%5Cgeqslant%20n)

Тогда

![](https://latex.codecogs.com/svg.image?xy=b%5Ekb%5Em,%20%5C:%5C:k&plus;m%5Cleq%20n,%20%5C:%5C:m%5Cneq%200%5C%5C%5Comega%20=b%5Ekb%5Emb%5E%7Bn-k-m%7Daaa%5En)

"Накачка" y:

![](https://latex.codecogs.com/svg.image?%5Comega%20=b%5Ek(b%5Em)%5Eib%5E%7Bn-k-m%7Daaa%5En%5C%5Ci=0%5C;%5C;%5CRightarrow%20%5C;%5C;%20%5Comega_0%20=b%5E%7Bn-m%7Daaa%5En%5C%5Cm%5Cneq%200%5C;%5C;%5CRightarrow%20%5Comega_0%5Cnotin%20L%20)

Следовательно, язык не является регулярным.



Язык № 3:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Ba%5Emw%20%5Cmid%20w%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%201%20%5Cleqslant%20%7Cw%7C_b%20%5Cleqslant%20m%5C%7D%20)

Лемма о накачке:

![](https://latex.codecogs.com/svg.image?L-%5Ctext%7Bregular%7D%5CRightarrow%20%5Cexists%20n%5C;%5C;%5C;%5Cforall%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cexists%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cforall%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cin%20L)

Отрицание:

![](https://latex.codecogs.com/svg.image?%5Cforall%20n%5C;%5C;%5C;%5Cexists%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cforall%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cexists%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cnotin%20%20L)

Возьмём слово

![](https://latex.codecogs.com/svg.image?%5Comega%20=%20a%5Enb%5En,%20%5Cforall%20n%5C%5C%5Cleft%7C%5Comega%20%20%5Cright%7C%20%5Cgeqslant%20n)

Тогда

![](https://latex.codecogs.com/svg.image?xy=a%5Eka%5Em,%20%5C:%5C:k&plus;m%5Cleq%20n,%20%5C:%5C:m%5Cneq%200%5C%5C%5Comega%20=a%5Eka%5Ema%5E%7Bn-k-m%7Db%5En)

"Накачка" y:

![](https://latex.codecogs.com/svg.image?%5Comega%20=a%5Ek(a%5Em)%5Eia%5E%7Bn-k-m%7Db%5En%5C%5Ci=0%5C;%5C;%5CRightarrow%20%5C;%5C;%20%5Comega_0%20=a%5E%7Bn-m%7Db%5En%5C%5Cm%5Cneq%200%5C;%5C;%5CRightarrow%20%5Comega_0%5Cnotin%20L)

Следовательно, язык не является регулярным.



Язык № 4:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Ba%5Ekb%5Ema%5En%20%5Cmid%20k%20=%20n%20%5Cvee%20m%20%3E%200%5C%7D)

Лемма о накачке:

![](https://latex.codecogs.com/svg.image?L-%5Ctext%7Bregular%7D%5CRightarrow%20%5Cexists%20n%5C;%5C;%5C;%5Cforall%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cexists%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cforall%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cin%20L)

Отрицание:

![](https://latex.codecogs.com/svg.image?%5Cforall%20n%5C;%5C;%5C;%5Cexists%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cforall%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cexists%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cnotin%20%20L)

Возьмём слово

![](https://latex.codecogs.com/svg.image?%5Comega%20=%20a%5Enba%5En,%20%5Cforall%20n%5C%5C%5Cleft%7C%5Comega%20%20%5Cright%7C%20%5Cgeqslant%20n)

Тогда

![](https://latex.codecogs.com/svg.image?xy=a%5Eka%5Em,%20%5C:%5C:k&plus;m%5Cleq%20n,%20%5C:%5C:m%5Cneq%200%5C%5C%5Comega%20=a%5Eka%5Ema%5E%7Bn-k-m%7Dba%5En)

"Накачка" y:

![](https://latex.codecogs.com/svg.image?%5Comega%20=a%5Ek(a%5Em)%5Eia%5E%7Bn-k-m%7Dba%5En%5C%5Ci=2%5C;%5C;%5CRightarrow%20%5C;%5C;%20%5Comega_2%20=a%5E%7Bn&plus;m%7Dba%5En%5C%5Cm%5Cneq%200%5C;%5C;%5CRightarrow%20%5Comega_2%5Cnotin%20L)

Следовательно, язык не является регулярным.



Язык № 5:

![](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bucv%20%5Cmid%20u%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20v%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20u%20%5Cneq%20v%5ER%5C%7D)

Лемма о накачке:

![](https://latex.codecogs.com/svg.image?L-%5Ctext%7Bregular%7D%5CRightarrow%20%5Cexists%20n%5C;%5C;%5C;%5Cforall%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cexists%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cforall%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cin%20L)

Отрицание:

![](https://latex.codecogs.com/svg.image?%5Cforall%20n%5C;%5C;%5C;%5Cexists%20%5Comega%20%5Cin%20L,%5C;%5C;%5Cleft%7C%5Comega%5Cright%7C%5Cgeqslant%20n%5C;%5C;%5C;%5Cforall%20x,y,z%5C;%5C;%5C;%5Comega%20=xyz%5C;%5C;%5C;%5Cleft%7Cxy%5Cright%7C%5Cleqslant%20n%5C;%5C;%5C;%5C%5Cy%5Cneq%20%5Cvarepsilon%20%5C;%5C;%5C;%5Cexists%20i%5Cgeqslant%200%5C;%5C;%5C;xy%5Eiz%5Cnotin%20%20L)

Возьмём слово

![](https://latex.codecogs.com/svg.image?%5Comega%20=%20(ab)%5Enc(ab)%5En%20=%20s_1s_2...s_n...s_%7B2n%7D...s_%7B4n%7Ds_%7B4n&plus;1%7D,%20%5Cforall%20n%5C%5C%5Cleft%7C%5Comega%20%20%5Cright%7C%20%5Cgeqslant%20n)

Тогда

![](https://latex.codecogs.com/svg.image?xy=(s_1s_2...s_k)(s_%7Bk&plus;1%7Ds_%7Bk&plus;2%7D...s_%7Bk&plus;m%7D),%20%5C:%5C:k&plus;m%5Cleq%20n,%20%5C:%5C:m%5Cneq%200%5C%5C%5Comega%20=(s_1s_2...s_k)(s_%7Bk&plus;1%7Ds_%7Bk&plus;2%7D...s_%7Bk&plus;m%7D)(s_%7Bk&plus;m&plus;1%7Ds_%7Bk&plus;m&plus;2%7D...s_%7B2n%7Dc(ab)%5En))

"Накачка" y:

![](https://latex.codecogs.com/svg.image?%5Comega%20=(s_1s_2...s_k)(s_%7Bk&plus;1%7Ds_%7Bk&plus;2%7D...s_%7Bk&plus;m%7D)%5Ei(s_%7Bk&plus;m&plus;1%7Ds_%7Bk&plus;m&plus;2%7D...s_%7B2n%7Dc(ab)%5En)%5C%5Ci=2%5C;%5C;%5CRightarrow%20%5C;%5C;%20%5Comega_2%20=(s_1s_2...s_k)(s_%7Bk&plus;1%7Ds_%7Bk&plus;2%7D...s_%7Bk&plus;m%7D)%5E2(s_%7Bk&plus;m&plus;1%7Ds_%7Bk&plus;m&plus;2%7D...s_%7B2n%7Dc(ab)%5En)%5C%5Cm%5Cneq%200%5C;%5C;%5CRightarrow%20%5Comega_2%5Cnotin%20L)

Следовательно, язык не является регулярным.