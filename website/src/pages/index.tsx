import type {ReactNode} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';
import BadgeCheck from '@site/static/img/badge-check.svg';

import './index.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero', 'heroBanner')}>
      <div className={clsx('container', 'heroBanner_container')}>
        <Heading as="h1" className="hero__title">
          <img src={String(siteConfig.themeConfig.banner)} alt={siteConfig.title} />
        </Heading>
        <p className="hero__title">{siteConfig.tagline}</p>
        <div className={clsx('heroFeatures')}>
          <p className="hero__subtitle">
            Features:
          </p>
          <ul>
            <li><BadgeCheck/> Commands – Get notified when a long-running command finishes.</li>
            <li><BadgeCheck/> Processes – Get notified when a background process exits.</li>
            {/* <li><BadgeCheck/> Dynamic Output – Detect changes in command output (great for waiting on conditions).</li> */}
          </ul>
        </div>
        <div className={clsx('buttons')}>
          {/* <Link
            className="button button--secondary button--lg"
            to="/docs/intro">
            Introduction
          </Link> */}
            <Link
              className="button button--primary button--lg"
              to="/docs/quickstart/overview">
              Quickstart
            </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={siteConfig.title}
      description="Description will go into a meta tag in <head />">
      <HomepageHeader />
      {/* <main>
        <HomepageFeatures />
      </main> */}
    </Layout>
  );
}
