import { Kafka } from "kafkajs";

class KafkaConfig {
  kafka;
  producer;
  consumer;
  constructor() {
    this.kafka = new Kafka({
      clientId: "nodejs-kafka",
      brokers: ["kafka:29092"],
    });
    this.producer = this.kafka.producer();
    this.consumer = this.kafka.consumer({ groupId: "payment-server" });
  }

  async produce(topic: any, messages: any) {
    try {
      await this.producer.connect();
      await this.producer.send({
        topic: topic,
        messages: messages,
      });
    } catch (error) {
      console.error(error);
    } finally {
      await this.producer.disconnect();
    }
  }

  async consume(topic: any, callback: any) {
    try {
      await this.consumer.connect();
      await this.consumer.subscribe({ topic: topic, fromBeginning: true });
      await this.consumer.run({
        eachMessage: async ({ topic, partition, message }: {topic: any, partition: any, message: any}) => {
          const value = message.value.toString();
          callback(value);
        },
      });
    } catch (error) {
      console.error(error);
    }
  }
}

export default KafkaConfig;
