PGDMP     5    /            
    |         
   article_db    15.2    15.2      +           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ,           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            -           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            .           1262    52383 
   article_db    DATABASE     �   CREATE DATABASE article_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE article_db;
                postgres    false                        3079    52398 	   uuid-ossp 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
    DROP EXTENSION "uuid-ossp";
                   false            /           0    0    EXTENSION "uuid-ossp"    COMMENT     W   COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
                        false    2            �            1259    52427    articles    TABLE     q  CREATE TABLE public.articles (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text NOT NULL,
    category text NOT NULL,
    content text NOT NULL,
    author_id uuid NOT NULL,
    status text DEFAULT 'draft'::text NOT NULL
);
    DROP TABLE public.articles;
       public         heap    postgres    false    2            �            1259    52442    comments    TABLE     (  CREATE TABLE public.comments (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    article_id uuid NOT NULL,
    content text NOT NULL,
    user_id uuid NOT NULL
);
    DROP TABLE public.comments;
       public         heap    postgres    false    2            �            1259    52385    roles    TABLE     N   CREATE TABLE public.roles (
    id bigint NOT NULL,
    name text NOT NULL
);
    DROP TABLE public.roles;
       public         heap    postgres    false            �            1259    52384    roles_id_seq    SEQUENCE     u   CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.roles_id_seq;
       public          postgres    false    216            0           0    0    roles_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;
          public          postgres    false    215            �            1259    52409    users    TABLE     ?  CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text NOT NULL,
    password text NOT NULL,
    email text NOT NULL,
    role_id bigint NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false    2            |           2604    52388    roles id    DEFAULT     d   ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);
 7   ALTER TABLE public.roles ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            '          0    52427    articles 
   TABLE DATA           w   COPY public.articles (id, created_at, updated_at, deleted_at, title, category, content, author_id, status) FROM stdin;
    public          postgres    false    218   d$       (          0    52442    comments 
   TABLE DATA           h   COPY public.comments (id, created_at, updated_at, deleted_at, article_id, content, user_id) FROM stdin;
    public          postgres    false    219   �%       %          0    52385    roles 
   TABLE DATA           )   COPY public.roles (id, name) FROM stdin;
    public          postgres    false    216   m&       &          0    52409    users 
   TABLE DATA           k   COPY public.users (id, created_at, updated_at, deleted_at, username, password, email, role_id) FROM stdin;
    public          postgres    false    217   �&       1           0    0    roles_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.roles_id_seq', 7, true);
          public          postgres    false    215            �           2606    52435    articles articles_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.articles DROP CONSTRAINT articles_pkey;
       public            postgres    false    218            �           2606    52449    comments comments_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.comments DROP CONSTRAINT comments_pkey;
       public            postgres    false    219            �           2606    52392    roles roles_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_pkey;
       public            postgres    false    216            �           2606    52394    roles uni_roles_name 
   CONSTRAINT     O   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uni_roles_name UNIQUE (name);
 >   ALTER TABLE ONLY public.roles DROP CONSTRAINT uni_roles_name;
       public            postgres    false    216            �           2606    52420    users uni_users_email 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT uni_users_email;
       public            postgres    false    217            �           2606    52418    users uni_users_username 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_username UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT uni_users_username;
       public            postgres    false    217            �           2606    52416    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    217            �           1259    52441    idx_articles_deleted_at    INDEX     R   CREATE INDEX idx_articles_deleted_at ON public.articles USING btree (deleted_at);
 +   DROP INDEX public.idx_articles_deleted_at;
       public            postgres    false    218            �           1259    52460    idx_comments_deleted_at    INDEX     R   CREATE INDEX idx_comments_deleted_at ON public.comments USING btree (deleted_at);
 +   DROP INDEX public.idx_comments_deleted_at;
       public            postgres    false    219            �           1259    52426    idx_users_deleted_at    INDEX     L   CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);
 (   DROP INDEX public.idx_users_deleted_at;
       public            postgres    false    217            �           2606    52436    articles fk_articles_author    FK CONSTRAINT     |   ALTER TABLE ONLY public.articles
    ADD CONSTRAINT fk_articles_author FOREIGN KEY (author_id) REFERENCES public.users(id);
 E   ALTER TABLE ONLY public.articles DROP CONSTRAINT fk_articles_author;
       public          postgres    false    3211    218    217            �           2606    52455    comments fk_comments_article    FK CONSTRAINT     �   ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_comments_article FOREIGN KEY (article_id) REFERENCES public.articles(id);
 F   ALTER TABLE ONLY public.comments DROP CONSTRAINT fk_comments_article;
       public          postgres    false    219    3213    218            �           2606    52450    comments fk_comments_user    FK CONSTRAINT     x   ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES public.users(id);
 C   ALTER TABLE ONLY public.comments DROP CONSTRAINT fk_comments_user;
       public          postgres    false    3211    219    217            �           2606    52421    users fk_users_role    FK CONSTRAINT     r   ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES public.roles(id);
 =   ALTER TABLE ONLY public.users DROP CONSTRAINT fk_users_role;
       public          postgres    false    3202    217    216            '   '  x���1k�0Fg�Wh/gN��$y�ҭ]�1�,�h�M�ђ_Ou
�}ox�c���L`�: �	1$$&6�3h��
��h�4D���	���{=���|j�yn��%]��?��ڧ�|<494u�����L+8t)2W�-pL���J�P��"1��weN��5�X|ɠ�&��6Q\�&9��\i48��E�
�׎wYt�=!��|�DOĜS�\�:Bp����*�Tx+ˣu��!j#e��cM�B�J�6@��) qa�[F���7޲]ז�}.����H�n����vؒ      (   �   x���1�1E�S�Gى��b+J�$��$��k�EHt+�����b/f�� �d�b+�sI�;�12R��Gݣ,�Q%�~�ǟ�;5t����!���	r��^t:�Ώ�\�v��u�k�N�};�__&cY#f�A��&4E��u3t7����kw�"�����Q�*,:�{(����z��?��y�� ah�      %   -   x�3�tL����2�-N-�2�t/M-.�2�t,-��/����� �!	�      &   �  x���Mo�0��3���M���9�Q�c#Ti�c;�a!)��~�4��]��~��/Sf!J�Ph(��h�3B+��70Ĵ~=@`�q�	�_ �O}56y�0�m��j!�*�F���Sy܇/q�cTI����|''��M���pqZ~ϒ�����ʏz�k�S뭗�jZ��\+$VP���' &�`0�uGW@`0�	�ct���5^Ue��b���`��Wo��u�>{�v��i�g�����[7��l���y�1(l*����?r�T����R�Z�-�� ���X)���Q0�qƥ����ӎ/v��PUd���~gҎ+'�Yz���s�ey�����b}���2r�7T8� 	�5�$�ԇS�h�1W<@ �1^��_?����	�&CJ���s%��ج�k�������S�r���Rn��N�.#W��k6��)��7     