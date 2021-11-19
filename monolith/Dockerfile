FROM ruby:2.7

RUN apt-get update -qq \
    && apt-get install -yq --no-install-recommends \
    build-essential \
    less \
    vim

ENV LANG=C.UTF-8 \
  GEM_HOME=/bundle \
  BUNDLE_JOBS=4 \
  BUNDLE_RETRY=3
ENV BUNDLE_PATH $GEM_HOME
ENV BUNDLE_APP_CONFIG=$BUNDLE_PATH \
  BUNDLE_BIN=$BUNDLE_PATH/bin
ENV PATH $BUNDLE_BIN:$PATH

RUN mkdir -p /app

WORKDIR /app

COPY . .
RUN bundle install --jobs 4 --retry 3

EXPOSE 9292
CMD ["bundle", "exec", "rackup", "-p", "9292", "--host", "0.0.0.0"]
STOPSIGNAL SIGTERM
