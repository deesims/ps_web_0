--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.2
-- Dumped by pg_dump version 9.6.2

SET statement_timeout = 0;
SET lock_timeout = 0;
-- SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS coopcat;
--
-- Name: coopcat; Type: DATABASE; Schema: -; Owner: coopcat_dev
--

CREATE DATABASE coopcat;


ALTER DATABASE coopcat OWNER TO coopcat_dev;

\connect coopcat

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: update_num_available_positions(); Type: FUNCTION; Schema: public; Owner: coopcat_dev
--

CREATE FUNCTION update_num_available_positions() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN	
		UPDATE job
		SET num_available_positions = (SELECT num_available_positions - 1
						FROM job
						WHERE JOB_ID = NEW.JOBJOB_ID)
		WHERE JOB_ID = NEW.JOBJOB_ID;
		
		RETURN NEW;
	END;
$$;


ALTER FUNCTION public.update_num_available_positions() OWNER TO coopcat_dev;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: company; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE company (
    company_id numeric(19,0) NOT NULL,
    name character varying(150) NOT NULL,
    description character varying(255) NOT NULL,
    num_employees integer NOT NULL
);


ALTER TABLE company OWNER TO coopcat_dev;

--
-- Name: company_company_id_seq; Type: SEQUENCE; Schema: public; Owner: coopcat_dev
--

CREATE SEQUENCE company_company_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE company_company_id_seq OWNER TO coopcat_dev;

--
-- Name: company_company_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: coopcat_dev
--

ALTER SEQUENCE company_company_id_seq OWNED BY company.company_id;


--
-- Name: job; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE job (
    job_id numeric(19,0) NOT NULL,
    company_id numeric(19,0) NOT NULL,
    name character varying(100),
    rating integer,
    student_level integer,
    num_of_positions integer NOT NULL,
    deadline_date date NOT NULL,
    num_available_positions integer
);


ALTER TABLE job OWNER TO coopcat_dev;

--
-- Name: job_job_id_seq; Type: SEQUENCE; Schema: public; Owner: coopcat_dev
--

CREATE SEQUENCE job_job_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE job_job_id_seq OWNER TO coopcat_dev;

--
-- Name: job_job_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: coopcat_dev
--

ALTER SEQUENCE job_job_id_seq OWNED BY job.job_id;


--
-- Name: resume; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE resume (
    resume_id numeric(19,0) NOT NULL,
    author_id numeric(19,0) NOT NULL,
    last_updated_at timestamp without time zone NOT NULL,
    resume_path character varying(200)
);


ALTER TABLE resume OWNER TO coopcat_dev;

--
-- Name: resume_Resume_ID_seq; Type: SEQUENCE; Schema: public; Owner: coopcat_dev
--

CREATE SEQUENCE "resume_Resume_ID_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "resume_Resume_ID_seq" OWNER TO coopcat_dev;

--
-- Name: resume_Resume_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: coopcat_dev
--

ALTER SEQUENCE "resume_Resume_ID_seq" OWNED BY resume.resume_id;


--
-- Name: resume_review; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE resume_review (
    resume_id numeric(19,0) NOT NULL,
    moderator_id numeric(19,0) NOT NULL,
    review_date timestamp without time zone NOT NULL,
    review character varying(300) NOT NULL
);


ALTER TABLE resume_review OWNER TO coopcat_dev;

--
-- Name: student; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE student (
    program character varying(200),
    school_start_date date,
    school_end_date date,
    student_level integer NOT NULL,
    user_id numeric(19,0) NOT NULL
);


ALTER TABLE student OWNER TO coopcat_dev;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: coopcat_dev
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO coopcat_dev;

--
-- Name: user; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE users (
    user_id numeric(19,0) DEFAULT nextval('user_id_seq'::regclass) NOT NULL,
    name character varying(100) NOT NULL,
    password character varying(250) NOT NULL,
    email character varying(100) NOT NULL,
    address character varying(100),
    role integer,
    CONSTRAINT "USER_email_check" CHECK (((email)::text ~* '^[A-Za-z0-9]+@[A-Za-z0-9]+.[A-Za-z0-9]+$'::text))
);


ALTER TABLE users OWNER TO coopcat_dev;

--
-- Name: work_review; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE work_review (
    comments character varying(300) NOT NULL,
    rating integer NOT NULL,
    works_forjobjob_id numeric(19,0) NOT NULL,
    works_for_user_id numeric(19,0) NOT NULL
);


ALTER TABLE work_review OWNER TO coopcat_dev;

--
-- Name: works_for; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE works_for (
    job_id numeric(19,0) NOT NULL,
    user_id numeric(19,0) NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL
);


ALTER TABLE works_for OWNER TO coopcat_dev;

--
-- Name: company company_id; Type: DEFAULT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY company ALTER COLUMN company_id SET DEFAULT nextval('company_company_id_seq'::regclass);


--
-- Name: job job_id; Type: DEFAULT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY job ALTER COLUMN job_id SET DEFAULT nextval('job_job_id_seq'::regclass);


--
-- Name: resume resume_id; Type: DEFAULT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY resume ALTER COLUMN resume_id SET DEFAULT nextval('"resume_Resume_ID_seq"'::regclass);


--
-- Data for Name: company; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Name: company_company_id_seq; Type: SEQUENCE SET; Schema: public; Owner: coopcat_dev
--

SELECT pg_catalog.setval('company_company_id_seq', 1, false);


--
-- Data for Name: job; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Name: job_job_id_seq; Type: SEQUENCE SET; Schema: public; Owner: coopcat_dev
--

SELECT pg_catalog.setval('job_job_id_seq', 1, false);


--
-- Data for Name: resume; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Name: resume_Resume_ID_seq; Type: SEQUENCE SET; Schema: public; Owner: coopcat_dev
--

SELECT pg_catalog.setval('"resume_Resume_ID_seq"', 1, false);


--
-- Data for Name: resume_review; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: student; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--

INSERT INTO users (user_id, name, password, email, address, role) VALUES (4, 'John Doe1', 'password', 'johndoe1@gmail.com', '123 North Street', 1);
INSERT INTO users (user_id, name, password, email, address, role) VALUES (6, 'John Doe2', 'password', 'johndoe2@gmail.com', '123 North Street', 1);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: coopcat_dev
--

SELECT pg_catalog.setval('user_id_seq', 6, true);


--
-- Data for Name: work_review; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: works_for; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Name: user USER_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY users
    ADD CONSTRAINT "USER_pkey" PRIMARY KEY (user_id);


--
-- Name: company company_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY company
    ADD CONSTRAINT company_pkey PRIMARY KEY (company_id);


--
-- Name: job job_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY job
    ADD CONSTRAINT job_pkey PRIMARY KEY (job_id);


--
-- Name: resume resume_author_id_resume_id_pk; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY resume
    ADD CONSTRAINT resume_author_id_resume_id_pk PRIMARY KEY (author_id, resume_id);


--
-- Name: work_review review_works_forjobjob_id_works_for_user_id_pk; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY work_review
    ADD CONSTRAINT review_works_forjobjob_id_works_for_user_id_pk PRIMARY KEY (works_forjobjob_id, works_for_user_id);


--
-- Name: resume_review reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY resume_review
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (resume_id, moderator_id);


--
-- Name: student student_user_id_pk; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY student
    ADD CONSTRAINT student_user_id_pk PRIMARY KEY (user_id);


--
-- Name: works_for works_for_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY works_for
    ADD CONSTRAINT works_for_pkey PRIMARY KEY (job_id, user_id);


--
-- Name: user_email_uindex; Type: INDEX; Schema: public; Owner: coopcat_dev
--

CREATE UNIQUE INDEX user_email_uindex ON users USING btree (email);


--
-- Name: works_for update_num_available_positions_trigger; Type: TRIGGER; Schema: public; Owner: coopcat_dev
--

CREATE TRIGGER update_num_available_positions_trigger AFTER INSERT ON works_for FOR EACH ROW EXECUTE PROCEDURE update_num_available_positions();


--
-- Name: resume_review fkreviews794469; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY resume_review
    ADD CONSTRAINT fkreviews794469 FOREIGN KEY (moderator_id) REFERENCES users(user_id);


--
-- Name: student fkstudent664699; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY student
    ADD CONSTRAINT fkstudent664699 FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE;


--
-- Name: job job_company_company_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY job
    ADD CONSTRAINT job_company_company_id_fk FOREIGN KEY (company_id) REFERENCES company(company_id);


--
-- Name: resume resume_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY resume
    ADD CONSTRAINT resume_user_id_fk FOREIGN KEY (author_id) REFERENCES users(user_id);


--
-- Name: works_for works_for_job_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY works_for
    ADD CONSTRAINT works_for_job_id_fk FOREIGN KEY (job_id) REFERENCES job(job_id);


--
-- Name: works_for works_for_user_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY works_for
    ADD CONSTRAINT works_for_user_user_id_fk FOREIGN KEY (user_id) REFERENCES users(user_id);


--
-- PostgreSQL database dump complete
--

