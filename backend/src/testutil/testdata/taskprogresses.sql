--
-- PostgreSQL database dump
--

-- Dumped from database version 12.15 (Debian 12.15-1.pgdg110+1)
-- Dumped by pg_dump version 14.9 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

SET CONSTRAINTS ALL DEFERRED;

--
-- Data for Name: Task; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public."tasks" (id, "task_id", "task_content", "task_category", "task_name") FROM stdin;
1	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	model1model1	DB設計	DBモデリング1
2	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	model2model2	DB設計	DBモデリング2
3	2f05b674-6e6d-4b37-b343-4e17953a4a33	unit1unit1	テスト	単体テスト1
4	a3b4f0ea-b188-403b-9141-97a044535972	unit2unit2	テスト	単体テスト2
5	75b71703-0b3d-4901-a91b-0d66e78f944f	tdd1tdd1	設計	TDD
\.

--
-- Data for Name: Team; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public."teams" (id, "team_id", "team_name") FROM stdin;
1	315b4766-bd70-4084-b91f-850f242953af	1
2	37194894-fee5-4ee1-9373-1f112bc7e401	2
\.


--
-- Data for Name: Pair; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public."pairs" (id, "pair_id", "pair_name", "team_id") FROM stdin;
1	413b1afb-9675-49ba-9615-5f03f29adc1c	A	315b4766-bd70-4084-b91f-850f242953af
2	545a3bec-3a5d-4047-b59c-aca8bb74fefb	B	315b4766-bd70-4084-b91f-850f242953af
3	c3d644f9-c8d5-45a0-8ca5-5d6dd20556b2	C	37194894-fee5-4ee1-9373-1f112bc7e401
\.

--
-- Data for Name: User; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public."users" (email, user_status, id, "user_id", "user_name", "pair_id", "team_id") FROM stdin;
test1@example.com	ACTIVE	1	d18d97ee-4034-40a3-99aa-95f4a8a40226	test1	413b1afb-9675-49ba-9615-5f03f29adc1c	315b4766-bd70-4084-b91f-850f242953af
test2@example.com	ACTIVE	2	9669a94f-6afb-447e-b76d-be37956d7dd2	test2	413b1afb-9675-49ba-9615-5f03f29adc1c	315b4766-bd70-4084-b91f-850f242953af
test3@example.com	ACTIVE	3	ffa0c412-157e-49e3-bfb2-5b71071fc003	test3	545a3bec-3a5d-4047-b59c-aca8bb74fefb	315b4766-bd70-4084-b91f-850f242953af
test4@example.com	ACTIVE	4	add2589e-fc47-4495-b093-4301ff52bee4	test4	545a3bec-3a5d-4047-b59c-aca8bb74fefb	315b4766-bd70-4084-b91f-850f242953af
test5@example.com	ACTIVE	5	f5d68717-abfc-4871-b2da-04d2468de055	test5	c3d644f9-c8d5-45a0-8ca5-5d6dd20556b2	37194894-fee5-4ee1-9373-1f112bc7e401
test6@example.com	ACTIVE	6	ef1a7553-1d46-4a1e-99df-3df79ff57663	test6	c3d644f9-c8d5-45a0-8ca5-5d6dd20556b2	37194894-fee5-4ee1-9373-1f112bc7e401
test7@example.com	ACTIVE	7	3af8efb9-e17a-4619-81bb-5a84d263ecbc	test7	c3d644f9-c8d5-45a0-8ca5-5d6dd20556b2	37194894-fee5-4ee1-9373-1f112bc7e401
test8@example.com	INACTIVE	8	c94f6203-d30f-40f2-888a-604088ca8144	test8	\N	\N
test9@example.com	DELETE	9	42f10819-f385-434e-8b37-9dcce775f885	test9	\N	\N
\.


--
-- Data for Name: TaskProgress; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public."task_progresses" (id, "task_id", "user_id", "task_progress_id", "task_status") FROM stdin;
1	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	d18d97ee-4034-40a3-99aa-95f4a8a40226	f5183782-623d-43af-85fe-92934898aa31	NOT_STARTED
2	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	d18d97ee-4034-40a3-99aa-95f4a8a40226	0ede5b95-3fea-4b89-8c24-013a32519a71	COMPLETED
3	2f05b674-6e6d-4b37-b343-4e17953a4a33	d18d97ee-4034-40a3-99aa-95f4a8a40226	3580cbed-d362-4e9e-93b1-22e3856c649a	COMPLETED
4	a3b4f0ea-b188-403b-9141-97a044535972	d18d97ee-4034-40a3-99aa-95f4a8a40226	4e6002a3-83f7-487f-aa8f-c0f79cfe3bc7	COMPLETED
5	75b71703-0b3d-4901-a91b-0d66e78f944f	d18d97ee-4034-40a3-99aa-95f4a8a40226	fa15696c-4560-4ed0-9258-4f305599c073	NOT_STARTED
6	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	9669a94f-6afb-447e-b76d-be37956d7dd2	15929d89-b269-4e25-b55f-64e946d170de	COMPLETED
7	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	9669a94f-6afb-447e-b76d-be37956d7dd2	49ccaa0a-1c91-41e3-8d59-5e440f4ee6ad	COMPLETED
8	2f05b674-6e6d-4b37-b343-4e17953a4a33	9669a94f-6afb-447e-b76d-be37956d7dd2	8a8a8f6c-97d6-4a85-9655-dca3fa779411	PROGRESS
9	a3b4f0ea-b188-403b-9141-97a044535972	9669a94f-6afb-447e-b76d-be37956d7dd2	ee8a09a4-ac50-4cae-b776-6a95b3a3bbcd	NOT_STARTED
10	75b71703-0b3d-4901-a91b-0d66e78f944f	9669a94f-6afb-447e-b76d-be37956d7dd2	ce00f4bc-87c9-4387-89f1-f5b8ed9bac8a	NOT_STARTED
11	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	ffa0c412-157e-49e3-bfb2-5b71071fc003	600b2306-a534-4281-aad6-0d40dbc89192	NOT_STARTED
12	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	ffa0c412-157e-49e3-bfb2-5b71071fc003	306583b6-d11e-4b63-9603-5f6e42d9fe5c	NOT_STARTED
13	2f05b674-6e6d-4b37-b343-4e17953a4a33	ffa0c412-157e-49e3-bfb2-5b71071fc003	ed7f1adc-6f55-4c28-b359-92d4e2da30d2	NOT_STARTED
14	a3b4f0ea-b188-403b-9141-97a044535972	ffa0c412-157e-49e3-bfb2-5b71071fc003	3694fcf3-67f9-46a8-8a2c-261dc01d994b	NOT_STARTED
15	75b71703-0b3d-4901-a91b-0d66e78f944f	ffa0c412-157e-49e3-bfb2-5b71071fc003	f5410376-b617-42d5-b1d5-eb5e36949bf0	NOT_STARTED
16	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	add2589e-fc47-4495-b093-4301ff52bee4	4e205367-565a-421d-b649-ce155bc214cc	COMPLETED
17	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	add2589e-fc47-4495-b093-4301ff52bee4	91f8d0f5-31f8-48af-b3ac-66c3710d2f49	COMPLETED
18	2f05b674-6e6d-4b37-b343-4e17953a4a33	add2589e-fc47-4495-b093-4301ff52bee4	f5d1226d-2533-4d5d-b6dc-87be26571480	COMPLETED
19	a3b4f0ea-b188-403b-9141-97a044535972	add2589e-fc47-4495-b093-4301ff52bee4	fcf922da-40d0-4210-ad3e-ffe9f9dd17e1	COMPLETED
20	75b71703-0b3d-4901-a91b-0d66e78f944f	add2589e-fc47-4495-b093-4301ff52bee4	d71074c8-f8da-43f0-a858-1a5a28f5c9f7	NOT_STARTED
21	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	f5d68717-abfc-4871-b2da-04d2468de055	a3732b55-d383-42bc-8e69-7ef7ed3456cb	COMPLETED
22	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	f5d68717-abfc-4871-b2da-04d2468de055	51ed4e73-fc52-47a8-b9de-e4832168a8e4	COMPLETED
23	2f05b674-6e6d-4b37-b343-4e17953a4a33	f5d68717-abfc-4871-b2da-04d2468de055	b4521238-b1d2-415f-86c6-00a8307d61cb	COMPLETED
24	a3b4f0ea-b188-403b-9141-97a044535972	f5d68717-abfc-4871-b2da-04d2468de055	f62ebaf2-6921-4c47-8980-4002fc895114	COMPLETED
25	75b71703-0b3d-4901-a91b-0d66e78f944f	f5d68717-abfc-4871-b2da-04d2468de055	3f493f95-54c3-4520-9238-387fba3c93e1	NOT_STARTED
26	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	ef1a7553-1d46-4a1e-99df-3df79ff57663	ae7d865c-beb6-4fa3-ba39-08d3396b7cdf	COMPLETED
27	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	ef1a7553-1d46-4a1e-99df-3df79ff57663	146b3295-a491-4c37-a53d-c534a7ea2ee6	COMPLETED
28	2f05b674-6e6d-4b37-b343-4e17953a4a33	ef1a7553-1d46-4a1e-99df-3df79ff57663	3ba5a672-28cc-4b27-896e-f399c166ca74	COMPLETED
29	a3b4f0ea-b188-403b-9141-97a044535972	ef1a7553-1d46-4a1e-99df-3df79ff57663	8a566faf-6384-4948-a525-e6c44e3b7a42	COMPLETED
30	75b71703-0b3d-4901-a91b-0d66e78f944f	ef1a7553-1d46-4a1e-99df-3df79ff57663	a3ac0913-7bb1-43a0-8879-f54d1ab1fb8a	NOT_STARTED
31	076c17e6-3f48-49a2-a4a1-68c3888ba2cd	3af8efb9-e17a-4619-81bb-5a84d263ecbc	8d99c274-9d7d-46c2-bd59-40830b14a52e	COMPLETED
32	d848ec90-fc48-49a5-b41c-6fa70d25a0aa	3af8efb9-e17a-4619-81bb-5a84d263ecbc	2bde0f81-96e4-4168-9a12-cb5740c0762d	COMPLETED
33	2f05b674-6e6d-4b37-b343-4e17953a4a33	3af8efb9-e17a-4619-81bb-5a84d263ecbc	1918d7eb-c41a-4af9-832a-54093bea0e94	COMPLETED
34	a3b4f0ea-b188-403b-9141-97a044535972	3af8efb9-e17a-4619-81bb-5a84d263ecbc	c4367f19-821b-436d-9fd4-083352274bfa	COMPLETED
\.


SET CONSTRAINTS ALL IMMEDIATE;