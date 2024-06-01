
<p align="center">
  <a href="https://www.spurtcms.com/#gh-light-mode-only">
    <img src="https://www.spurtcms.com/spurtcms.png" width="318px" alt="Spurtcms logo" />
  </a>
   
</p>
<h3 align="center">Open Source Golang based CMS Solution - Self hosted </h3>
<p align="center"> Build with Golang + PostgreSQL</p>

<br />
<p align="center">
  <a href="https://github.com/spurtcms/spurtcms-admin/releases">
    <img src="https://img.shields.io/github/last-commit/spurtcms/deployment" alt="GitHub last commit" />
  </a>
  <a href="https://github.com/spurtcms/spurtcms-admin/issues">
    <img src="https://img.shields.io/github/issues/spurtcms/deployment" alt="GitHub issues" />
  </a>

  <a href="https://github.com/spurtcms/spurtcms-admin/releases">
    <img src="https://img.shields.io/github/repo-size/spurtcms/deployment?color=orange" alt="GitHub repo size" />
  </a>
</p>
<br />

> [!IMPORTANT]
> 🎉 <strong>Spurtcms 1.0 is now available!</strong> Read more in the <a target="_blank" href="https://www.spurtcms.com/spurtcms-change-log" rel="dofollow"><strong>announcement post</strong></a>.
<br />
<p>
### Welcome to the World of Modern Web Application Development in Go
At spurtCMS, where we're on a mission to redefine Web Application Development for the modern era. Inspired by the legacy of traditional CMS platforms like WordPress and Joomla, SpurtCMS offers developers a new approach to building dynamic web experiences on a modern Golang architecture. spurtCMS leverages a variety of advanced technologies and features to deliver a comprehensive web development solution, including a GraphQL API that offers single endpoint access.
</p>
<br />
### Our Vision

Our vision is to revolutionize content management and beyond by providing developers with a platform that combines the best aspects of traditional CMS platforms with the power and flexibility of modern Golang architecture. With our modular and decoupled design, SpurtCMS allows developers to extract and use any package for any project, giving them unparalleled flexibility and control.

### Our Mission
Our mission is to empower developers to create exceptional websites, applications, and services with ease. Whether you're building a learning management system (LMS), a blog, an e-commerce platform, or any other type of application, SpurtCMS offers the tools and flexibility you need to bring your vision to life. With our go templates, you can use our packages to enhance existing projects or build end-to-end solutions from scratch.

### Our Goals
At SpurtCMS, our goals are simple: to empower developers, drive innovation, and create value for our users. We believe in pushing the boundaries of what's possible in web development, and we're committed to providing developers with the tools they need to succeed.

### Our Team
Our team is made up of passionate developers, designers, and innovators who share a common goal: to build the best platform in the world for web development. With years of experience in Golang and a shared commitment to excellence, we're dedicated to making SpurtCMS the go-to solution for developers everywhere.

### Why SpurtCMS?

SpurtCMS offers developers a unique combination of power and flexibility. With our modular architecture and decoupled design, developers can extract and use any package for any project, whether it's an LMS, a blog, an e-commerce platform, or something else entirely. Our go templates simplify the integration of our packages, whether you're enhancing existing projects or developing complete solutions.
Join us on this journey as we redefine the future of web development and empower developers to build the applications and services of tomorrow with SpurtCMS.






## ❯  🚀 Easy to Deploy Spurtcms Admin Panel on your server

This is the official repository of Spurtcms. Using these Build , you can easily deploy spurtcms in your local server.

### Step 1: Download the source files:

Clone the Git repository that contains spurtCMS Admin project files, PostgreSQL dump file and .env file from the path https://github.com/spurtcms-admin.git using the “git clone” command.

```
https://github.com/spurtcms/spurtcms-admin.git
```
After successful git clone, you should see a folder “spurtcms-admin” with folders locales, view, storage, public and files such as binary file, spurtCMS-admin.sql file and  .env, file.


### Step 2: Database Setup

Utilize the "Restore" feature in PgAdmin to populate the database with the necessary content from the database dump spurtCMS-admin.sql cloned in the above step.

Locate the .env file inside the project folder “spurtcms-admin-app” and configure it with the details of newly created database such as database name, user name, password etc

#### PostgreSQL Database Configuration

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=your_database_name
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_SSL_MODE=disable
```

Successful completion of this step completes the database configuration for spurtCMS Admin application.


### Step 3: Running the Project

Open the terminal within the project / cloned folder “spurtcms-admin-app”, note down the binary file name and execute the following command:

```
./{binary-file-name}
```
This command initiates the spurtCMS Admin application, allowing you to begin your journey with this powerful content management system.

 

By following the steps outlined in this article, you have successfully set up spurtCMS Admin on your system. Ensure that all prerequisites are met and the configuration steps are accurately executed to enjoy a seamless experience with spurtCMS Admin application. Now you can explore the features and functionalities of spurtCMS Admin for efficient content management.


live demo of our intuitive Admin Panel .

```
Username : Admin
Password : Admin@123
```


## 🤔 Support , Document and Help

spurtcms 4.8.2 is published to npm under the `@spurtcms/*` namespace.

You can find our extended documentation on our [https://spurtcms.com/documentation](https://spurtcms.com/documentation), but some quick links that might be helpful:

- Read [Technology](https://www.spurtcms.com/opensource-ecommerce-multivendor-nodejs-react-angular) to learn about our vision and what's in the box.

- Our [Discard](https://discord.com/invite/9TNgqUY24N) Questions, Live Discussions [spurtcms Support](https://picco.support).

- Some [Video](https://www.youtube.com/@spurtcms/videos) Video Tutorials 
- Every [Release](https://github.com/spurtcms/spurtcms-admin/releases) is documented on the Github Releases page.

🐞 If you spot a bug, please [submit a detailed issue](https://github.com/spurtcms/spurtcms-admin/issues/new), and wait for assistance.




## ❯ Maintainers
spurtcms is developed and maintain by [Piccosoft Software Labs India (P) Limited,](https://www.piccosoft.com).


## ❯ License

spurtcms is released under the [BSD-3-Clause License.](https://github.com/spurtcms/spurtcms/blob/master/LICENSE).



