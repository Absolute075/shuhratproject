import { useEffect, useMemo, useState } from 'react';

type Slide = {
  src: string;
  srcSet?: string;
  sizes?: string;
  alt: string;
};

export default function HeroSlider() {
  const slides: Slide[] = useMemo(
    () => [
      {
        src: 'https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng.jpg',
        srcSet:
          'https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng-p-500.jpg 500w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng-p-800.jpg 800w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng-p-1080.jpg 1080w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng-p-1600.jpg 1600w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng-p-2000.jpg 2000w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng-p-2600.jpg 2600w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245dc9d14b6757a5dd0834_eng.jpg 2940w',
        sizes: '(max-width: 2940px) 100vw, 2940px',
        alt: ''
      },
      {
        src: 'https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245ee98b2b92ad3f406dfe_MGE.jpg',
        srcSet:
          'https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245ee98b2b92ad3f406dfe_MGE-p-500.jpg 500w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245ee98b2b92ad3f406dfe_MGE-p-800.jpg 800w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245ee98b2b92ad3f406dfe_MGE-p-1080.jpg 1080w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245ee98b2b92ad3f406dfe_MGE-p-1600.jpg 1600w, https://cdn.prod.website-files.com/6924575d25f2c55da5fad02c/69245ee98b2b92ad3f406dfe_MGE.jpg 1920w',
        sizes: '(max-width: 1920px) 100vw, 1920px',
        alt: ''
      }
    ],
    []
  );

  const [active, setActive] = useState(0);

  useEffect(() => {
    const id = window.setInterval(() => {
      setActive((v) => (v + 1) % slides.length);
    }, 6000);
    return () => window.clearInterval(id);
  }, [slides.length]);

  return (
    <div className="slider-section">
      <div className="slider w-slider">
        <div className="w-slider-mask pp-slider">
          {slides.map((s, idx) => (
            <div key={s.src} className={`pp-slide${idx === active ? ' is-active' : ''}`}>
              <img
                sizes={s.sizes}
                srcSet={s.srcSet}
                alt={s.alt}
                src={s.src}
                loading="lazy"
              />
            </div>
          ))}
        </div>

        <div className="left-arrow w-slider-arrow-left" onClick={() => setActive((v) => (v - 1 + slides.length) % slides.length)}>
          <div className="icon w-icon-slider-left"></div>
        </div>
        <div className="right-arrow w-slider-arrow-right" onClick={() => setActive((v) => (v + 1) % slides.length)}>
          <div className="w-icon-slider-right"></div>
        </div>

        <div className="slide-nav w-slider-nav w-round"></div>
      </div>
    </div>
  );
}
