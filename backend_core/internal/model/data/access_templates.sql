--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.3

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
SET session_replication_role = 'replica';


--
-- Data for Name: access_templates; Type: TABLE DATA; Schema: public; Owner: eunimartuser
--

INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (1,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'SUPER_USER');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (2,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'COMPANY_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (3,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'MDM_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (4,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'ORDERS_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (5,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'INVENTORY_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (6,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'INVENTORY_TASKS_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (7,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'ACCOUNTING_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (8,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'RETURNS_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (9,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'SHIPPING_ADMIN');
INSERT INTO public.access_templates( id, is_enabled, is_active, created_by, updated_date, updated_by, deleted_by, created_date, company_id, app_id, deleted_at, name) VALUES (10,true, true, 1, '2022-06-21 16:11:58.270482+05:30', 1, NULL, '2022-06-21 16:11:58.270482+05:30', 1, NULL, NULL, 'OMNICHANNEL_ADMIN');

SET session_replication_role = 'origin';


--
-- Name: access_templates_id_seq; Type: SEQUENCE SET; Schema: public; Owner: eunimartuser
--

select setval( pg_get_serial_sequence('public.access_templates', 'id'), (select max(id) from public.access_templates));

--
-- PostgreSQL database dump complete
--
