# go-knight

Knight's tour - make a pawn visit all tiles on a chequerboard; in golang.

A knight's tour is a sequence of moves of a knight on a chessboard such that the knight visits every square only once. If the knight ends on a square that is one knight's move from the beginning square (so that it could tour the board again immediately, following the same path), the tour is closed, otherwise it is open.

The problem is solved using H C Warnsdorff's technique which does the following:

Start from any tile and mark it as visited.
- To decide next tile in path, look at all possible unmarked tiles based on moving rules.
- Rank each possibility by the number of next moves from that tile.
- Move to any tile with the lowest rank.
- Add chosen tile to knight's tour path (i.e marked) and repeat the process from last chosen tile.

# Execution

Start by cloning the repository
- go inside the directory `cd go-knight`
- set starting positions by setting env variable `START_X` and `START_Y`
- run `START_X=2 START_Y=3 go run cmd/knights_tour/main.go`

# Output


```
....
2019/08/21 15:29:31 points visited:  97
22	48	45	23	71	58	12	68	59	13	
55	1	20	56	16	19	61	15	18	62	
46	24	91	47	44	92	70	41	11	69	
21	49	54	26	72	57	17	67	60	14	
86	2	43	87	90	42	78	89	35	63	
51	25	94	50	53	93	73	40	10	66	
6	82	85	27	79	88	34	64	75	31	
0	3	52	98	95	37	77	96	36	39	
84	28	7	83	29	8	74	30	9	65	
5	81	0	4	80	97	33	38	76	32	
2019/08/21 15:29:31 points visited:  98
22	48	45	23	71	58	12	68	59	13	
55	1	20	56	16	19	61	15	18	62	
46	24	91	47	44	92	70	41	11	69	
21	49	54	26	72	57	17	67	60	14	
86	2	43	87	90	42	78	89	35	63	
51	25	94	50	53	93	73	40	10	66	
6	82	85	27	79	88	34	64	75	31	
99	3	52	98	95	37	77	96	36	39	
84	28	7	83	29	8	74	30	9	65	
5	81	0	4	80	97	33	38	76	32	
2019/08/21 15:29:31 points visited:  99
22	48	45	23	71	58	12	68	59	13	
55	1	20	56	16	19	61	15	18	62	
46	24	91	47	44	92	70	41	11	69	
21	49	54	26	72	57	17	67	60	14	
86	2	43	87	90	42	78	89	35	63	
51	25	94	50	53	93	73	40	10	66	
6	82	85	27	79	88	34	64	75	31	
99	3	52	98	95	37	77	96	36	39	
84	28	7	83	29	8	74	30	9	65	
5	81	100	4	80	97	33	38	76	32	
2019/08/21 15:29:31 points visited:  100
```
